
all: fmt vet test

fmt:
	@go fmt ./...

vet:
	@go vet ./...

test:
	@go test -race ./...

codegen:
	@go run tools/codegen/main.go

.PHONY: all fmt vet test codegen