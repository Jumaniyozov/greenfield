GO_VERSION :=1.23.4

.PHONY: install-go init-go build test coverage report check-format install-lint static-check install
setup: install-go init-go

install: install-go init-go install-lint


install-go:
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

init-go:
	echo 'export PATH=$$PATH:/usr/local/go/bin' >> $${HOME}/.bashrc
	echo 'export PATH=$$PATH:$${HOME}/go/bin' >> $${HOME}/.bashrc

install-lint:
	sudo curl -sSfL \
	https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh\
	| sh -s -- -b $$(go env GOPATH)/bin v1.62.2

check-format:
	test -z "$$(go fmt ./...)"

static-check:
	golangci-lint run

test:
	go test ./... -coverprofile=coverage.out

check: check-format static-check test

coverage:
	go tool cover -func coverage.out | grep "total:" | awk '{print ((int($$3) > 80) != 1) }'

report:
	go tool cover -html=coverage.out -o cover.html

build:
	go build -o api cmd/main.go