
all: fmt vet test

fmt:
	@go fmt ./...

vet:
	@go vet ./...

test:
	@go test -race -cover ./...

codegen:
	@go run tools/codegen/main.go
	make fmt

.PHONY: all fmt vet test codegen
