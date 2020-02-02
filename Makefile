export GO111MODULE := on
export APP_DOTENV_PATH := $(shell pwd)/.env

GOOS := linux
GOARCH := amd64

go:
	gofmt -s -w .

run:
	go run cmd/main.go

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build cmd/api/main.go

test:
	go test -v ./...

dev-deps:
	GO111MODULE=off go get -u -v \
		github.com/oxequa/realize

refresh-run:
	realize start
