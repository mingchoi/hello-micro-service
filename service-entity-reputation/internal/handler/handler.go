package handler

import (
	"context"
	"fmt"

	proto "github.com/mingchoi/hello-micro-service/proto-go/proto-entity-reputation"
)

type ServiceHandler struct {
}

func (s *ServiceHandler) Get(ctx context.Context, req *proto.GetRequest) (*proto.GetResponse, error) {
	fmt.Println("Yoooo" + req.UserId)
	return &proto.GetResponse{
		Reputation: &proto.Reputation{UserName: "Youre gay"},
	}, nil
}

func (s *ServiceHandler) GetList(context.Context, *proto.GetListRequest) (*proto.GetListResponse, error) {
	return &proto.GetListResponse{}, nil
}

func (s *ServiceHandler) Create(context.Context, *proto.CreateRequest) (*proto.CreateResponse, error) {
	return &proto.CreateResponse{}, nil
}

func (s *ServiceHandler) Update(context.Context, *proto.UpdateRequest) (*proto.UpdateResponse, error) {
	return &proto.UpdateResponse{}, nil
}

func (s *ServiceHandler) Delete(context.Context, *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	return &proto.DeleteResponse{}, nil
}
