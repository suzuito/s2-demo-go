GO_SOURCES := $(shell find . -name '*.go')

api.exe:
	go build -o api.exe main_api/*.go

start-api:
	source dev.sh && $(shell go env GOPATH)/bin/air -c .air-api.toml