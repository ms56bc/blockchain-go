# Makefile

proto_files = $(wildcard proto/*.proto)
go_out_path = .

all: compile

compile:
	protoc $(proto_files) \
		--go_out=$(go_out_path) \
		--go-grpc_out=$(go_out_path) \
#	--go_opt=paths=source_relative \
#		--go-grpc_opt=paths=source_relative \
		--proto_path=.

clean:
	rm -f p2p/*.pb.go p2p/*.pb.gw.go

.PHONY: all compile clean bootstrap


bootstrap:
	nohup go run cmd/seed_node.go > seed_node.log 2>&1 & \
    nohup go run cmd/peer1.go > peer1.log 2>&1 & \
    nohup go run cmd/peer2.go > peer2.log 2>&1 &