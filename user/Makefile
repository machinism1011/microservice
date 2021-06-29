
GOPATH:=$(shell go env GOPATH)
.PHONY: init
init:
	go get -u github.com/golang/protobuf/proto
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get github.com/micro/micro/v2/cmd/protoc-gen-micro
	go get github.com/micro/micro/v2/cmd/protoc-gen-openapi

.PHONY: proto
proto:
	docker run --rm -v "$(shell pwd):$(shell pwd)" -w "$(shell pwd)" ericnothing0113/go-protoc:1.0.0 --proto_path=. --micro_out=. --go_out=. -I. proto/user/*.proto
	
.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: dockerBuild
dockerBuild:
	docker build -t user:latest .
