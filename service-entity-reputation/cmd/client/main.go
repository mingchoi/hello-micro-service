package main

import (
	"context"
	"fmt"
	"log"

	proto "github.com/mingchoi/hello-micro-service/proto-go/proto-entity-reputation"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := proto.NewReputationEntityServiceClient(conn)
	r, err := c.Get(context.TODO(), &proto.GetRequest{
		UserId: "YourMom",
	})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println(r)
}
