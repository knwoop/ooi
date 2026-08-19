BIN      ?= ooi
VERSION  ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS  := -s -w -X github.com/knwoop/ooi/cmd.version=$(VERSION)

.PHONY: build build-internal test lint clean

## build: plain binary (users provide ~/.config/ooi/credentials.json)
build:
	go build -ldflags "$(LDFLAGS)" -o $(BIN) .

## build-internal: binary with the OAuth client embedded, for handing out to
## internal users so they only need `ooi auth` && `ooi install`.
##   OOI_CLIENT_ID=... OOI_CLIENT_SECRET=... make build-internal
build-internal:
	@test -n "$(OOI_CLIENT_ID)"     || (echo "OOI_CLIENT_ID is required"     >&2; exit 1)
	@test -n "$(OOI_CLIENT_SECRET)" || (echo "OOI_CLIENT_SECRET is required" >&2; exit 1)
	go build -ldflags "$(LDFLAGS) \
		-X github.com/knwoop/ooi/internal/calendar.embeddedClientID=$(OOI_CLIENT_ID) \
		-X github.com/knwoop/ooi/internal/calendar.embeddedClientSecret=$(OOI_CLIENT_SECRET)" \
		-o $(BIN) .

test:
	go vet ./...
	go test -race ./...

lint:
	golangci-lint run ./...

clean:
	rm -f $(BIN)
