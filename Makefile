.PHONY: test

test:
	@go test -v ./...

.PHONY: build
build:
	go build -o glox-cli ./cmd/main.go