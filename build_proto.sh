#!/usr/bin/sh
protoc --go_out=plugins=grpc:person person/person.proto
protoc --go_out=plugins=grpc:$GOPATH/src chat/chat.proto
