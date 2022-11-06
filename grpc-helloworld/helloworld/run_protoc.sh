#!/bin/bash


echo
echo "to install protoc plugins:"
echo "go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28"
echo "go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2"
echo
echo
echo


protoc	--go_out=.  \
	--go-grpc_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	helloworld.proto

