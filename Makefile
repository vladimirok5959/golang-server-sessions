VERSION="1.0.5"

default: test

test:
	@go test -v ./...
