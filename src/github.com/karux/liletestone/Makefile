# vi: ft=make

GOPATH:=$(shell go env GOPATH)

.PHONY: proto test

proto:
#	go get github.com/golang/protobuf/protoc-gen-go
	protoc -I . ./model/liletestone.proto --lile-server_out=. --go_out=plugins=grpc:${GOPATH}/src

client:
	protoc -I . ./model/liletestone.proto --cobra_out=plugins=client:${GOPATH}/src

build: proto
#	go build -o build/karux/liletestone karux/liletestone/liletestone/main.go
	go build -i -o ${GOBIN}/liletestone ./liletestone/main.go
#	go build -i -o ${GOBIN}/liletestoneClient ./liletestone.cobra.pb.go

test:
	@go get github.com/rakyll/gotest
	gotest -p 1 -v ./...
