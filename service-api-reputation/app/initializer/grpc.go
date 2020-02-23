package initializer

import (
	"log"

	proto "github.com/mingchoi/hello-micro-service/proto-go/proto-entity-reputation"
	"github.com/mingchoi/hello-micro-service/service-api-reputation/app"
	"google.golang.org/grpc"
)

func InitGRPC() *grpc.ClientConn {
	conn, err := grpc.Dial("service-entity-reputation:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	app.EntityReputation = proto.NewReputationEntityServiceClient(conn)
	return conn
}
