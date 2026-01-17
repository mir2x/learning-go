.DEFAULT_GOAL := build
.PHONY: clean fmt vet build

clean:
	rm hello_world

fmt:
	go fmt ./...

vet: fmt
	go vet ./...

build: vet
	go build
	
