GO_SOURCES := $(shell find . -name '*.go')

api.exe: ${GO_SOURCES}
	go build -o api.exe main_api/*.go

data.exe: ${GO_SOURCES}
	go build -o data.exe cmd/data/*.go

*/*/result.txt: ${GO_SOURCES}
	FILE_PATH_GEOJSON=$(dir $@)/result.geojson go run $(dir $@)/main.go | tee $@

clean:
	rm *.exe