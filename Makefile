GO_SOURCES := $(shell find . -name '*.go')

api.exe: ${GO_SOURCES}
	go build -o api.exe main_api/*.go

start-api:
	source dev.sh && $(shell go env GOPATH)/bin/air -c .air-api.toml

*/*/result.txt: ${GO_SOURCES}
	go run $(dir $@)/main.go | tee $@

*/*/result.geojson: ${GO_SOURCES}
	go run $(dir $@)/main.go | tee $@