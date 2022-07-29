package main

import (
	grpcListener "go-grpc/pkg/grpc-listener"
)

func main() {

	grpc := grpcListener.GRPCListener{}

	grpc.Listen()

}
