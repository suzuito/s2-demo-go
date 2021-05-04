GO_SOURCES := $(shell find . -name '*.go')

api.exe: ${GO_SOURCES}
	go build -o api.exe main_api/*.go

data.exe: ${GO_SOURCES}
	go build -o data.exe cmd/data/*.go

build-cmd-article:
	docker-compose exec -T env bash -c "make clean && make data.exe"

data-build: build-cmd-article
	docker-compose exec -T env bash -c "./data.exe build -input-article-dir ${DIR}"
	./execute-sample-codes.sh ${DIR}
	docker-compose exec -T env bash -c "./data.exe build-index -input-dir $(dir ${DIR})"

*/*/*/*/result.txt: ${GO_SOURCES}
	docker-compose exec -T env bash -c "DIR_PATH_RESULT=$(dir $@) go run $(dir $@)main.go | tee $@"

rsync-godzilla: build-cmd-article
	docker-compose exec -T env bash -c "./data.exe upload -input-dir ${DIR}"

clean:
	rm *.exe