
all: fmt vet test

fmt:
	@go fmt ./...

vet:
	@go vet ./...

test:
	@go test -race -cover ./...

codegen:
	rm -rf tulipindicators
	git clone https://github.com/TulipCharts/tulipindicators.git
	(cd tulipindicators && git checkout cffa15e && rm -rf .git)
	@go run tools/codegen/main.go
	make fmt

.PHONY: all fmt vet test codegen
