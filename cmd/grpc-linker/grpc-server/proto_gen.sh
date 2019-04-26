#!/bin/bash
protoc -I proto --go_out=plugins=grpc:proto proto/apiauth.proto
#protoc -I proto --php_out=plugins=grpc:proto proto/apiauth.proto