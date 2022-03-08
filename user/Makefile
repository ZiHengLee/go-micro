
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v3/cmd/protoc-gen-micro
	go get github.com/micro/micro/v3/cmd/protoc-gen-openapi

.PHONY: api
api:
	protoc --openapi_out=. --proto_path=. proto/user.proto

.PHONY: proto
proto:
    sudo docker run --rm -v $(shell pwd):$(shell pwd) -w $(shell pwd)  znly/protoc -I ./ --go_out=./ --micro_out=./  ./proto/user/*.proto
	protoc --proto_path=. --micro_out=. --go_out=:. proto/user.proto
	
.PHONY: build
build:
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user *.go


.PHONY: test
test:
	go test -v ./... -cover

.PHONY: dockerBuild
docker:
	docker build . -t user:latest
