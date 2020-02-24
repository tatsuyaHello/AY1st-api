export GO111MODULE := on
export APP_DOTENV_PATH := $(shell pwd)/.env

GOOS := linux
GOARCH := amd64

go:
	gofmt -s -w .

run:
	go run main.go

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build main.go

test:
	go test -v ./...

dev-deps:
	GO111MODULE=off go get -u -v \
		github.com/oxequa/realize

refresh-run:
	realize start

docker-up:
	@if [ ! $(docker-compose ps | grep mysql) ] ; then docker-compose up -d ; fi

init-db-local: docker-up
	DATABASE_NAME=ay1st-local sh ./fixtures/init_db_local.sh

