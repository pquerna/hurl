
all: test install
	@echo "Done"

install:
	go install github.com/pquerna/hurl/hurl

deps:
	go get -u launchpad.net/gocheck
	go get -u github.com/spf13/cobra
	go get -u github.com/dchest/uniuri

fmt:
	go fmt ./common
	go fmt ./http
	go fmt ./hurl
	go fmt ./ui
	go fmt ./workers

test: clean
	go test -v ./...

clean:
	go clean

.PHONY: deps clean test fmt install all
