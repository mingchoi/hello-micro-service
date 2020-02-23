package handler

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	proto "github.com/mingchoi/hello-micro-service/proto-go/proto-entity-reputation"
	"github.com/mingchoi/hello-micro-service/service-api-reputation/app"
)

func Hello(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "hi",
	})
}

func Handle(ctx *gin.Context) {
	r, err := app.EntityReputation.Get(context.TODO(), &proto.GetRequest{
		UserId: "YourMom",
	})
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	ctx.JSON(200, gin.H{
		"message": r.Reputation.UserName,
	})
}
