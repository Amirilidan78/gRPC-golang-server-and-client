package grpcListener

import (
	"fmt"
	grpc_servers "go-grpc/pkg/grpc-servers"
	"go-grpc/pkg/logs"
	"google.golang.org/grpc"
	"log"
	"net"
)

const grpcPort = "8000"

type GRPCListener struct {
}

func (g *GRPCListener) Listen() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()

	logs.RegisterLogServiceServer(s, &grpc_servers.LogServer{})

	log.Printf("gRPC server started on port %s", grpcPort)

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
