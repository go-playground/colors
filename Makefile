GOCMD=GO111MODULE=on go

lint:
	golangci-lint run --timeout=5m

test:
	$(GOCMD) test -cover -race ./...

bench:
	$(GOCMD) test -bench=. -benchmem ./...

.PHONY: test lint linters-install