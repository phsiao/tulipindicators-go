
all: fmt vet test

fmt:
	@go fmt ./...

vet: build
	@go vet ./...

test: build
	@go test -race -cover ./...

build:
	(cd tulipindicators && make)

codegen: build
	@go run tools/codegen/main.go
	make fmt

clean:
	(cd tulipindicators && make clean && rm -f rm example1 example2 fuzzer sample smoke)

.PHONY: all fmt vet test codegen build clean
