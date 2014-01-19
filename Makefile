
all: test
	go install github.com/pquerna/hurl/hurl
	@echo "Done"

deps:
	go get -u launchpad.net/gocheck
	go get -u github.com/spf13/cobra

fmt:
	go fmt ./hurl
	go fmt ./ui

test: clean
	go test -v ./...

clean:
	go clean

.PHONY: deps clean test format install
