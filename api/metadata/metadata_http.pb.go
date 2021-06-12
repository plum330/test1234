// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.0-rc3

package metadata

import (
	context "context"
	middleware "github.com/go-kratos/kratos/v2/middleware"
	transport "github.com/go-kratos/kratos/v2/transport"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = new(middleware.Middleware)
var _ = new(transport.Transporter)
var _ = binding.BindVars

const _ = http.SupportPackageIsVersion1

type MetadataHTTPServer interface {
	GetServiceDesc(context.Context, *GetServiceDescRequest) (*GetServiceDescReply, error)
	ListServices(context.Context, *ListServicesRequest) (*ListServicesReply, error)
}

func RegisterMetadataHTTPServer(s *http.Server, srv MetadataHTTPServer) {
	r := s.Route("/")
	r.GET("/services", _Metadata_ListServices0_HTTP_Handler(srv))
	r.GET("/services/{name}", _Metadata_GetServiceDesc0_HTTP_Handler(srv))
}

func _Metadata_ListServices0_HTTP_Handler(srv MetadataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListServicesRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		transport.SetOperation(ctx, "/kratos.api.Metadata/ListServices")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListServices(ctx, req.(*ListServicesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListServicesReply)
		return ctx.Result(200, reply)
	}
}

func _Metadata_GetServiceDesc0_HTTP_Handler(srv MetadataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetServiceDescRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		transport.SetOperation(ctx, "/kratos.api.Metadata/GetServiceDesc")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetServiceDesc(ctx, req.(*GetServiceDescRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetServiceDescReply)
		return ctx.Result(200, reply)
	}
}

type MetadataHTTPClient interface {
	GetServiceDesc(ctx context.Context, req *GetServiceDescRequest, opts ...http.CallOption) (rsp *GetServiceDescReply, err error)
	ListServices(ctx context.Context, req *ListServicesRequest, opts ...http.CallOption) (rsp *ListServicesReply, err error)
}

type MetadataHTTPClientImpl struct {
	cc *http.Client
}

func NewMetadataHTTPClient(client *http.Client) MetadataHTTPClient {
	return &MetadataHTTPClientImpl{client}
}

func (c *MetadataHTTPClientImpl) GetServiceDesc(ctx context.Context, in *GetServiceDescRequest, opts ...http.CallOption) (*GetServiceDescReply, error) {
	var out GetServiceDescReply
	path := binding.EncodeVars("/services/{name}", in, true)
	opts = append(opts, http.Operation("/kratos.api.Metadata/GetServiceDesc"))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MetadataHTTPClientImpl) ListServices(ctx context.Context, in *ListServicesRequest, opts ...http.CallOption) (*ListServicesReply, error) {
	var out ListServicesReply
	path := binding.EncodeVars("/services", in, true)
	opts = append(opts, http.Operation("/kratos.api.Metadata/ListServices"))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
