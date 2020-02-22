package main

import (
	"log"
	"net"

	proto "github.com/mingchoi/hello-micro-service/proto-go/proto-entity-reputation"
	"github.com/mingchoi/hello-micro-service/service-entity-reputation/app/handler"
	"github.com/mingchoi/hello-micro-service/service-entity-reputation/app/initializer"
	"google.golang.org/grpc"
)

func main() {
	initializer.InitXORM()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	proto.RegisterReputationEntityServiceServer(s, &handler.ServiceHandler{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
