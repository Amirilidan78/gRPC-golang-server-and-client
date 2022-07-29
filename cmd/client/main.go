package main

import (
	grpcClient "go-grpc/pkg/grpc-client"
)

func main() {

	grpc := grpcClient.GRPCClient{}

	grpc.LogViaGRPC()

}
