// Code generated by protoc-gen-go. DO NOT EDIT.
// source: reputation_service.proto

package proto_entity_reputation

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("reputation_service.proto", fileDescriptor_1a473637eb7d89b8) }

var fileDescriptor_1a473637eb7d89b8 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0xd1, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0xc6, 0xf1, 0xc3, 0xc2, 0x0a, 0xb9, 0x19, 0x0f, 0xca, 0x82, 0xe2, 0x7a, 0x6f, 0x0a, 0xfa,
	0x06, 0xba, 0x52, 0x11, 0x05, 0x59, 0xf1, 0xe2, 0x45, 0x76, 0xe3, 0xd0, 0x0e, 0xb4, 0x99, 0x98,
	0x4c, 0x05, 0x5f, 0xc6, 0x67, 0x15, 0x9b, 0xc6, 0x08, 0xbb, 0xcd, 0xad, 0xf4, 0xfb, 0xe7, 0x77,
	0x19, 0x71, 0xe2, 0xc0, 0xf6, 0xbc, 0x61, 0x24, 0xf3, 0xe6, 0xc1, 0x7d, 0xa2, 0x06, 0x65, 0x1d,
	0x31, 0xc9, 0x43, 0x30, 0x8c, 0xfc, 0xa5, 0x52, 0xb0, 0x38, 0xfa, 0x17, 0x3b, 0x17, 0xba, 0xcb,
	0xef, 0x99, 0x38, 0x5e, 0xff, 0xfd, 0xbf, 0x1d, 0x1e, 0x3d, 0x07, 0x49, 0xae, 0xc4, 0xac, 0x02,
	0x96, 0xa7, 0x6a, 0xc7, 0x52, 0x15, 0xf0, 0x1a, 0x3e, 0x7a, 0xf0, 0xbc, 0x38, 0x9b, 0x9a, 0xbd,
	0x25, 0xe3, 0x41, 0x3e, 0x89, 0x83, 0x0a, 0xf8, 0x01, 0x3d, 0xcb, 0xe5, 0xfe, 0xf4, 0x77, 0x8b,
	0xda, 0x45, 0x2e, 0x19, 0xc5, 0x47, 0x31, 0xbf, 0x71, 0xb0, 0x61, 0x90, 0xe7, 0x7b, 0xea, 0x30,
	0x45, 0x6f, 0x99, 0x29, 0x12, 0xf7, 0x62, 0xdf, 0xa7, 0xb8, 0x30, 0xe5, 0xb8, 0x58, 0x24, 0x6e,
	0x05, 0x2d, 0x4c, 0x70, 0x61, 0xca, 0x71, 0xb1, 0x08, 0xdc, 0xf5, 0xfd, 0xeb, 0x5d, 0x8d, 0xdc,
	0xf4, 0x5b, 0xa5, 0xa9, 0x2b, 0x3b, 0x34, 0xb5, 0x6e, 0x08, 0xcb, 0x06, 0xda, 0x96, 0x8a, 0x0e,
	0xb5, 0xa3, 0x62, 0xbc, 0x7c, 0x39, 0x5c, 0xb4, 0xa8, 0x69, 0xfc, 0x08, 0x74, 0x91, 0xe8, 0xed,
	0x7c, 0x18, 0xae, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xba, 0xe1, 0xce, 0xe0, 0x37, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReputationEntityServiceClient is the client API for ReputationEntityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReputationEntityServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type reputationEntityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReputationEntityServiceClient(cc grpc.ClientConnInterface) ReputationEntityServiceClient {
	return &reputationEntityServiceClient{cc}
}

func (c *reputationEntityServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/entity.reputation.ReputationEntityService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reputationEntityServiceClient) GetList(ctx context.Context, in *GetListRequest, opts ...grpc.CallOption) (*GetListResponse, error) {
	out := new(GetListResponse)
	err := c.cc.Invoke(ctx, "/entity.reputation.ReputationEntityService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reputationEntityServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/entity.reputation.ReputationEntityService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reputationEntityServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/entity.reputation.ReputationEntityService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reputationEntityServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/entity.reputation.ReputationEntityService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReputationEntityServiceServer is the server API for ReputationEntityService service.
type ReputationEntityServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
}

// UnimplementedReputationEntityServiceServer can be embedded to have forward compatible implementations.
type UnimplementedReputationEntityServiceServer struct {
}

func (*UnimplementedReputationEntityServiceServer) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedReputationEntityServiceServer) GetList(ctx context.Context, req *GetListRequest) (*GetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (*UnimplementedReputationEntityServiceServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedReputationEntityServiceServer) Update(ctx context.Context, req *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (*UnimplementedReputationEntityServiceServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterReputationEntityServiceServer(s *grpc.Server, srv ReputationEntityServiceServer) {
	s.RegisterService(&_ReputationEntityService_serviceDesc, srv)
}

func _ReputationEntityService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReputationEntityServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.reputation.ReputationEntityService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReputationEntityServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReputationEntityService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReputationEntityServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.reputation.ReputationEntityService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReputationEntityServiceServer).GetList(ctx, req.(*GetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReputationEntityService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReputationEntityServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.reputation.ReputationEntityService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReputationEntityServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReputationEntityService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReputationEntityServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.reputation.ReputationEntityService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReputationEntityServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReputationEntityService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReputationEntityServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/entity.reputation.ReputationEntityService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReputationEntityServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReputationEntityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "entity.reputation.ReputationEntityService",
	HandlerType: (*ReputationEntityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ReputationEntityService_Get_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _ReputationEntityService_GetList_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ReputationEntityService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ReputationEntityService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ReputationEntityService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reputation_service.proto",
}
