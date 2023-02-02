// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: to_video.proto

package to_video

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for ToVideoService service

func NewToVideoServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ToVideoService service

type ToVideoService interface {
	GetFavoriteStatus(ctx context.Context, in *GetFavoriteStatus_Request, opts ...client.CallOption) (*GetFavoriteStatus_Response, error)
}

type toVideoService struct {
	c    client.Client
	name string
}

func NewToVideoService(name string, c client.Client) ToVideoService {
	return &toVideoService{
		c:    c,
		name: name,
	}
}

func (c *toVideoService) GetFavoriteStatus(ctx context.Context, in *GetFavoriteStatus_Request, opts ...client.CallOption) (*GetFavoriteStatus_Response, error) {
	req := c.c.NewRequest(c.name, "ToVideoService.GetFavoriteStatus", in)
	out := new(GetFavoriteStatus_Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ToVideoService service

type ToVideoServiceHandler interface {
	GetFavoriteStatus(context.Context, *GetFavoriteStatus_Request, *GetFavoriteStatus_Response) error
}

func RegisterToVideoServiceHandler(s server.Server, hdlr ToVideoServiceHandler, opts ...server.HandlerOption) error {
	type toVideoService interface {
		GetFavoriteStatus(ctx context.Context, in *GetFavoriteStatus_Request, out *GetFavoriteStatus_Response) error
	}
	type ToVideoService struct {
		toVideoService
	}
	h := &toVideoServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ToVideoService{h}, opts...))
}

type toVideoServiceHandler struct {
	ToVideoServiceHandler
}

func (h *toVideoServiceHandler) GetFavoriteStatus(ctx context.Context, in *GetFavoriteStatus_Request, out *GetFavoriteStatus_Response) error {
	return h.ToVideoServiceHandler.GetFavoriteStatus(ctx, in, out)
}
