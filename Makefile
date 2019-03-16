VERSION="1.0.3"

default: test

test:
	@go test -v ./...
