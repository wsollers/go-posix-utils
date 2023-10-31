.PHONY: clean all

go-posix-utils: main.go
	go build -o bin/go-posix-utils main.go

clean: 
	rm -rf bin/*

all: clean go-posix-utils

run: go-posix-utils
	./bin/go-posix-utils

build-image: all
	docker build -t go-posix-utils .

run-image: build-image
	docker run  -p 4444:8080 --rm go-posix-utils

run-docker-scout: build-image
	docker scout cves go-posix-utils
