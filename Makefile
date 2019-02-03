VERSION="1.0.1"

default: test

test:
	@go test -v ./...
