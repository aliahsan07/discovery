package api

//go:generate protoc -I=. -I=$GOPATH/src --gogo_out=plugins=grpc:. repository-discovery-service.proto
