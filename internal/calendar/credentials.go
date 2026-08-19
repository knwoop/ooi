package calendar

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

// OAuth client credentials embedded at build time for internal distribution
// (see `make build-internal`):
//
//	go build -ldflags "-X github.com/knwoop/ooi/internal/calendar.embeddedClientID=... \
//	                   -X github.com/knwoop/ooi/internal/calendar.embeddedClientSecret=..."
//
// A "Desktop app" OAuth client secret is not treated as confidential by
// Google, so embedding it in a binary handed to internal users is acceptable.
// A credentials.json in the config directory always takes precedence, so
// binaries built without these values keep working as before.
var (
	embeddedClientID     string
	embeddedClientSecret string
)

// HasEmbeddedCredentials reports whether OAuth client credentials were
// embedded at build time.
func HasEmbeddedCredentials() bool {
	return embeddedClientID != "" && embeddedClientSecret != ""
}

// GetOAuthConfig returns the OAuth config, preferring credentials.json in the
// config directory and falling back to credentials embedded at build time.
func GetOAuthConfig() (*oauth2.Config, error) {
	configDir, err := ConfigDir()
	if err != nil {
		return nil, err
	}

	credentialsPath := filepath.Join(configDir, "credentials.json")
	b, err := os.ReadFile(credentialsPath)
	switch {
	case err == nil:
		config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
		if err != nil {
			return nil, fmt.Errorf("failed to parse credentials.json: %w", err)
		}
		return config, nil
	case errors.Is(err, os.ErrNotExist):
		if HasEmbeddedCredentials() {
			return &oauth2.Config{
				ClientID:     embeddedClientID,
				ClientSecret: embeddedClientSecret,
				Endpoint:     google.Endpoint,
				Scopes:       []string{calendar.CalendarReadonlyScope},
			}, nil
		}
		return nil, fmt.Errorf("no OAuth client credentials: place credentials.json in %s or use a build with embedded credentials", configDir)
	default:
		return nil, fmt.Errorf("failed to read credentials.json: %w", err)
	}
}
