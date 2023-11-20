.PHONY: clean all run build-image run-image run-docker-scout

build:  
	go build -o bin/go-posix-utils main.go 

go-posix-utils: build

clean: 
	rm -rf bin/*

all: clean build 

run: go-posix-utils 
	./bin/go-posix-utils

build-image: all
	docker build -t go-posix-utils .

run-image: build-image
	docker run  -p 4444:8080 --rm go-posix-utils

run-docker-scout: build-image
	docker scout cves go-posix-utils

ping-cnn: build
	sudo ./bin/go-posix-utils ping www.cnn.com 4000
