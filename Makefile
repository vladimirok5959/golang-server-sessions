VERSION="1.0.2"

default: test

test:
	@go test -v ./...
