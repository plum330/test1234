// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.0-rc3

package v1

import (
	context "context"
	transport "github.com/go-kratos/kratos/v2/transport"
	http "github.com/go-kratos/kratos/v2/transport/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = new(transport.Transporter)

const _ = http.SupportPackageIsVersion1

type MessageServiceHTTPServer interface {
	GetUserMessage(context.Context, *GetUserMessageRequest) (*GetUserMessageReply, error)
}

func RegisterMessageServiceHTTPServer(s *http.Server, srv MessageServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/message/user/{id}/{count}", _MessageService_GetUserMessage0_HTTP_Handler(srv))
}

func _MessageService_GetUserMessage0_HTTP_Handler(srv MessageServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetUserMessageRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		transport.SetOperation(ctx, "/api.message.v1.MessageService/GetUserMessage")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUserMessage(ctx, req.(*GetUserMessageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetUserMessageReply)
		return ctx.Result(200, reply)
	}
}

type MessageServiceHTTPClient interface {
	GetUserMessage(ctx context.Context, req *GetUserMessageRequest, opts ...http.CallOption) (rsp *GetUserMessageReply, err error)
}

type MessageServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewMessageServiceHTTPClient(client *http.Client) MessageServiceHTTPClient {
	return &MessageServiceHTTPClientImpl{client}
}

func (c *MessageServiceHTTPClientImpl) GetUserMessage(ctx context.Context, in *GetUserMessageRequest, opts ...http.CallOption) (*GetUserMessageReply, error) {
	var out GetUserMessageReply
	path := binding.EncodeVars("/v1/message/user/{id}/{count}", in, true)
	opts = append(opts, http.Operation("/api.message.v1.MessageService/GetUserMessage"))
	err := c.cc.Invoke(ctx, "GET", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
