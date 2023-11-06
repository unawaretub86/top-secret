# run test coverage
test:
	go test -v -cover ./...
# run linter run it to avoid waste actions minuts
lint:
	golangci-lint run