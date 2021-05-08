
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
	(cd tulipindicators && make clean && rm -f rm example1 example2 fuzzer sample smoke)
	make fmt

.PHONY: all fmt vet test codegen
