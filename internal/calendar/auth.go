package calendar

import (
	"context"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"

	"golang.org/x/oauth2"
)

const callbackPort = 8085

func Authenticate(ctx context.Context) (*oauth2.Token, error) {
	config, err := GetOAuthConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to get OAuth config: %w", err)
	}

	config.RedirectURL = fmt.Sprintf("http://localhost:%d/callback", callbackPort)

	codeCh := make(chan string, 1)
	errCh := make(chan error, 1)

	server := &http.Server{Addr: fmt.Sprintf(":%d", callbackPort)}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code == "" {
			errCh <- fmt.Errorf("no code in callback")
			http.Error(w, "No code received", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, `<html><body><h1>Authentication successful!</h1><p>You can close this window.</p></body></html>`)

		codeCh <- code
	})

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Opening browser for authentication...\n")
	if err := openBrowser(authURL); err != nil {
		fmt.Printf("Please open this URL manually:\n%s\n", authURL)
	}

	var code string
	select {
	case code = <-codeCh:
	case err := <-errCh:
		server.Shutdown(ctx)
		return nil, err
	case <-ctx.Done():
		server.Shutdown(ctx)
		return nil, ctx.Err()
	}

	server.Shutdown(ctx)

	token, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code for token: %w", err)
	}

	return token, nil
}

func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return fmt.Errorf("unsupported platform")
	}
	return cmd.Start()
}
