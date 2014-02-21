
all: test install
	@echo "Done"

install:
	go install github.com/pquerna/hurl/hurl

deps:
	go get -u launchpad.net/gocheck
	go get -u github.com/spf13/cobra
	go get -u github.com/dchest/uniuri
	go get -u github.com/cheggaaa/pb
	go get -u github.com/rcrowley/go-metrics
	go get -u github.com/coreos/go-etcd/etcd

fmt:
	go fmt github.com/pquerna/hurl/...

test: clean
	go test -v github.com/pquerna/hurl/...

clean:
	go clean -i github.com/pquerna/hurl/...

.PHONY: deps clean test fmt install all
