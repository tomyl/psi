
.PHONY: all test lint

test:
	go test -v

all: lint test

lint:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
