
GOPATH:=$(shell go env GOPATH)
.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=:. proto/crud/crud.proto && \
	protoc --proto_path=. --micro_out=. --go_out=:. proto/msg/msg.proto
	
.PHONY: build
build:
	go build -o ./bin/server ./cmd/server/... && go build -o ./bin/client ./cmd/client/...

.PHONY: client
client:
	go build -o ./bin/client ./cmd/client/...

.PHONY: server
server:
	go build -o ./bin/server ./cmd/server/...

.PHONY: run-server
run-server:
	./bin/server --broker=kafka --broker_address=localhost:9092

.PHONY: run-client
run-client:
	./bin/client --broker=kafka --broker_address=localhost:9092


.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t grpcProject:latest
