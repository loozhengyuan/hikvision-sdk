init:
	chmod -R +x .githooks
	git config core.hooksPath .githooks

install:
	go install -v

get:
	go get -v -u ./...

mod:
	go mod download
	go mod verify

fmt:
	go fmt ./...

lint:
	test -z $$(gofmt -l .)

vet:
	go vet -v ./...

test:
	go test -v -race -cover -coverprofile=coverage.txt -covermode=atomic ./...

docs:
	go doc -all

build: mod
	CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	go build \
		-ldflags="-X main.version=$$(git rev-parse --verify HEAD)" \
		-v -o main

clean: clean-mod clean-cov

clean-mod:
	go clean -modcache

clean-cov:
	find . -type f -name 'coverage.txt' -exec rm -rf {} \;
