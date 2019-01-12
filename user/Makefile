.PHONY: proto build docker
GOPATH:=$(shell go env GOPATH)

proto:
	protoc --proto_path=${GOPATH}/src/cinema/proto --go_out=pb/ --micro_out=pb/  ${GOPATH}/src/cinema/proto/user/*.proto


build: proto
	go build -o user-srv main.go plugin.go


docker:
	docker build . -t user-srv:latest
