package grpcClient

import (
	"context"
	"fmt"
	"go-grpc/pkg/logs"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

type GRPCClient struct {
}

func (g *GRPCClient) LogViaGRPC() {

	fmt.Println(time.Now().UnixMilli())

	conn, err := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	c := logs.NewLogServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.WriteLog(ctx, &logs.LogRequest{
		LogEntry: &logs.Log{
			Name: "name",
			Data: "data",
		},
	})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(time.Now().UnixMilli())
	log.Println(resp)

}
