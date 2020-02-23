package initializer

import (
	"log"
	"net"
	"strconv"

	proto "github.com/mingchoi/hello-micro-service/proto-go/proto-entity-reputation"
	"github.com/mingchoi/hello-micro-service/service-entity-reputation/app/handler"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func InitGRPC() {
	grpcPort := ":" + strconv.Itoa(viper.Get("app.grpc.port").(int))

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterReputationEntityServiceServer(s, &handler.ServiceHandler{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
