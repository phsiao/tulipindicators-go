
all: fmt vet test

fmt:
	@go fmt ./...

vet:
	@go vet ./...

test:
	@go test -race ./...

codegen:
	(cd tulipindicators && make)
	@go run tools/codegen/main.go
	make fmt

.PHONY: all fmt vet test codegen
