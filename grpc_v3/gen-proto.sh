#! /bin/bash

protoc -I mathproto/ --go_out=plugins=grpc:mathproto mathproto/math.proto