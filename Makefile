VERSION="1.0.0"

default: test

test:
	@go test -v ./...
