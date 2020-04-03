
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src/example --micro_out=. --go_out=. proto/example/example.proto

.PHONY: build
build: proto

	go build -o example-srv *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t example-srv:latest
