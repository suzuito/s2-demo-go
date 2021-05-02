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
	docker-compose exec -T env bash -c "FILE_PATH_GEOJSON=$(dir $@)/result.geojson go run $(dir $@)/main.go | tee $@"

rsync-godzilla: build-cmd-article
	# gsutil rsync -r ./data/article gs://suzuito-godzilla-s2-demo-article
	docker-compose exec -T env bash -c "./data.exe upload -input-dir ${DIR}"

clean:
	rm *.exe