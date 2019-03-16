VERSION="1.0.4"

default: test

test:
	@go test -v ./...
