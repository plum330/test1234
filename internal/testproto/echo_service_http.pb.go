// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.0-rc3

package testproto

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

type EchoServiceHTTPServer interface {
	Echo(context.Context, *SimpleMessage) (*SimpleMessage, error)
	EchoBody(context.Context, *SimpleMessage) (*SimpleMessage, error)
	EchoDelete(context.Context, *SimpleMessage) (*SimpleMessage, error)
	EchoPatch(context.Context, *DynamicMessageUpdate) (*DynamicMessageUpdate, error)
	EchoResponseBody(context.Context, *DynamicMessageUpdate) (*DynamicMessageUpdate, error)
}

func RegisterEchoServiceHTTPServer(s *http.Server, srv EchoServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/example/echo/{id}/{num}", _EchoService_Echo0_HTTP_Handler(srv))
	r.GET("/v1/example/echo/{id}/{num}/{lang}", _EchoService_Echo1_HTTP_Handler(srv))
	r.GET("/v1/example/echo1/{id}/{line_num}/{status.note}", _EchoService_Echo2_HTTP_Handler(srv))
	r.GET("/v1/example/echo2/{no.note}", _EchoService_Echo3_HTTP_Handler(srv))
	r.POST("/v1/example/echo/{id}", _EchoService_Echo4_HTTP_Handler(srv))
	r.POST("/v1/example/echo_body", _EchoService_EchoBody0_HTTP_Handler(srv))
	r.POST("/v1/example/echo_response_body", _EchoService_EchoResponseBody0_HTTP_Handler(srv))
	r.DELETE("/v1/example/echo_delete/{id}/{num}", _EchoService_EchoDelete0_HTTP_Handler(srv))
	r.PATCH("/v1/example/echo_patch", _EchoService_EchoPatch0_HTTP_Handler(srv))
}

func _EchoService_Echo0_HTTP_Handler(srv EchoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SimpleMessage
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		transport.SetOperation(ctx, "/testproto.EchoService/Echo")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Echo(ctx, req.(*SimpleMessage))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SimpleMessage)
		return ctx.Result(200, reply)
	}
}

func _EchoService_Echo1_HTTP_Handler(srv EchoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SimpleMessage
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		transport.SetOperation(ctx, "/testproto.EchoService/Echo")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Echo(ctx, req.(*SimpleMessage))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SimpleMessage)
		return ctx.Result(200, reply)
	}
}

func _EchoService_Echo2_HTTP_Handler(srv EchoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SimpleMessage
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		transport.SetOperation(ctx, "/testproto.EchoService/Echo")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Echo(ctx, req.(*SimpleMessage))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SimpleMessage)
		return ctx.Result(200, reply)
	}
}

func _EchoService_Echo3_HTTP_Handler(srv EchoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SimpleMessage
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		transport.SetOperation(ctx, "/testproto.EchoService/Echo")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Echo(ctx, req.(*SimpleMessage))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SimpleMessage)
		return ctx.Result(200, reply)
	}
}

func _EchoService_Echo4_HTTP_Handler(srv EchoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SimpleMessage
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		transport.SetOperation(ctx, "/testproto.EchoService/Echo")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Echo(ctx, req.(*SimpleMessage))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SimpleMessage)
		return ctx.Result(200, reply)
	}
}

func _EchoService_EchoBody0_HTTP_Handler(srv EchoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SimpleMessage
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		transport.SetOperation(ctx, "/testproto.EchoService/EchoBody")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.EchoBody(ctx, req.(*SimpleMessage))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SimpleMessage)
		return ctx.Result(200, reply)
	}
}

func _EchoService_EchoResponseBody0_HTTP_Handler(srv EchoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DynamicMessageUpdate
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		transport.SetOperation(ctx, "/testproto.EchoService/EchoResponseBody")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.EchoResponseBody(ctx, req.(*DynamicMessageUpdate))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DynamicMessageUpdate)
		return ctx.Result(200, reply.Body)
	}
}

func _EchoService_EchoDelete0_HTTP_Handler(srv EchoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SimpleMessage
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		transport.SetOperation(ctx, "/testproto.EchoService/EchoDelete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.EchoDelete(ctx, req.(*SimpleMessage))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SimpleMessage)
		return ctx.Result(200, reply)
	}
}

func _EchoService_EchoPatch0_HTTP_Handler(srv EchoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DynamicMessageUpdate
		if err := ctx.Bind(&in.Body); err != nil {
			return err
		}
		transport.SetOperation(ctx, "/testproto.EchoService/EchoPatch")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.EchoPatch(ctx, req.(*DynamicMessageUpdate))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DynamicMessageUpdate)
		return ctx.Result(200, reply)
	}
}

type EchoServiceHTTPClient interface {
	Echo(ctx context.Context, req *SimpleMessage, opts ...http.CallOption) (rsp *SimpleMessage, err error)
	EchoBody(ctx context.Context, req *SimpleMessage, opts ...http.CallOption) (rsp *SimpleMessage, err error)
	EchoDelete(ctx context.Context, req *SimpleMessage, opts ...http.CallOption) (rsp *SimpleMessage, err error)
	EchoPatch(ctx context.Context, req *DynamicMessageUpdate, opts ...http.CallOption) (rsp *DynamicMessageUpdate, err error)
	EchoResponseBody(ctx context.Context, req *DynamicMessageUpdate, opts ...http.CallOption) (rsp *DynamicMessageUpdate, err error)
}

type EchoServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewEchoServiceHTTPClient(client *http.Client) EchoServiceHTTPClient {
	return &EchoServiceHTTPClientImpl{client}
}

func (c *EchoServiceHTTPClientImpl) Echo(ctx context.Context, in *SimpleMessage, opts ...http.CallOption) (*SimpleMessage, error) {
	var out SimpleMessage
	path := binding.EncodeVars("/v1/example/echo/{id}", in, false)
	opts = append(opts, http.Operation("/testproto.EchoService/Echo"))
	err := c.cc.Invoke(ctx, "POST", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *EchoServiceHTTPClientImpl) EchoBody(ctx context.Context, in *SimpleMessage, opts ...http.CallOption) (*SimpleMessage, error) {
	var out SimpleMessage
	path := binding.EncodeVars("/v1/example/echo_body", in, false)
	opts = append(opts, http.Operation("/testproto.EchoService/EchoBody"))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *EchoServiceHTTPClientImpl) EchoDelete(ctx context.Context, in *SimpleMessage, opts ...http.CallOption) (*SimpleMessage, error) {
	var out SimpleMessage
	path := binding.EncodeVars("/v1/example/echo_delete/{id}/{num}", in, false)
	opts = append(opts, http.Operation("/testproto.EchoService/EchoDelete"))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *EchoServiceHTTPClientImpl) EchoPatch(ctx context.Context, in *DynamicMessageUpdate, opts ...http.CallOption) (*DynamicMessageUpdate, error) {
	var out DynamicMessageUpdate
	path := binding.EncodeVars("/v1/example/echo_patch", in, false)
	opts = append(opts, http.Operation("/testproto.EchoService/EchoPatch"))
	err := c.cc.Invoke(ctx, "PATCH", path, in.Body, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *EchoServiceHTTPClientImpl) EchoResponseBody(ctx context.Context, in *DynamicMessageUpdate, opts ...http.CallOption) (*DynamicMessageUpdate, error) {
	var out DynamicMessageUpdate
	path := binding.EncodeVars("/v1/example/echo_response_body", in, false)
	opts = append(opts, http.Operation("/testproto.EchoService/EchoResponseBody"))
	err := c.cc.Invoke(ctx, "POST", path, in, &out.Body, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
