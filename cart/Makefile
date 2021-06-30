
GOPATH:=$(shell go env GOPATH)
MODIFY=Mproto/imports/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    
	docker run --rm -v "$(shell pwd):$(shell pwd)" -w "$(shell pwd)" ericnothing0113/go-protoc:1.0.0 --proto_path=. --micro_out=. --go_out=. -I. proto/cart/*.proto
    

.PHONY: build
build: proto

	go build -o cart-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: dockerBuild
docker:
	docker build . -t cart-service:latest
