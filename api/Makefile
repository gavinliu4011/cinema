
GOPATH:=$(shell go env GOPATH)


.PHONY: proto
proto:
	protoc --proto_path=${GOPATH}/src/cinema/proto \
	 --micro_out=proto/ --go_out=proto/ ${GOPATH}/src/cinema/proto/user/*.proto

.PHONY: build
build: proto

	go build -o api-api main.go plugin.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t api-api:latest
