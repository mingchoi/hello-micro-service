package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	proto "github.com/mingchoi/hello-micro-service/proto-go/proto-entity-reputation"
	"google.golang.org/grpc"
)

func hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hi",
	})
}
func handle(ctx *gin.Context) {
	conn, err := grpc.Dial("service-entity-reputation:50051", grpc.WithInsecure(), grpc.WithBlock())
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

	ctx.JSON(200, gin.H{
		"message": r.Reputation.UserName,
	})
}

func main() {
	r := gin.Default()
	r.GET("/", hello)
	r.GET("/call", handle)
	r.Run(":8000")
}
