package grpc

import (
	"context"
	"github.com/go-kratos/kratos/v2/internal/matcher"

	"google.golang.org/grpc"
	grpcmd "google.golang.org/grpc/metadata"

	ic "github.com/go-kratos/kratos/v2/internal/context"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

// unaryServerInterceptor is a gRPC unary server interceptor
func (s *Server) unaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		ctx, cancel := ic.Merge(ctx, s.baseCtx)
		defer cancel()
		md, _ := grpcmd.FromIncomingContext(ctx)
		replyHeader := grpcmd.MD{}
		tr := &Transport{
			operation:   info.FullMethod,
			reqHeader:   headerCarrier(md),
			replyHeader: headerCarrier(replyHeader),
		}
		if s.endpoint != nil {
			tr.endpoint = s.endpoint.String()
		}
		ctx = transport.NewServerContext(ctx, tr)
		if s.timeout > 0 {
			ctx, cancel = context.WithTimeout(ctx, s.timeout)
			defer cancel()
		}
		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			return handler(ctx, req)
		}
		if next := s.middleware.Match(tr.Operation()); len(next) > 0 {
			h = middleware.Chain(next...)(h)
		}
		reply, err := h(ctx, req)
		if len(replyHeader) > 0 {
			_ = grpc.SetHeader(ctx, replyHeader)
		}
		return reply, err
	}
}

// wrappedStream is rewrite grpc stream's context
type wrappedStream struct {
	grpc.ServerStream
	ctx        context.Context
	middleware matcher.Matcher
}

func NewWrappedStream(ctx context.Context, stream grpc.ServerStream, m matcher.Matcher) grpc.ServerStream {
	return &wrappedStream{
		ServerStream: stream,
		ctx:          ctx,
		middleware:   m,
	}
}

func (w *wrappedStream) Context() context.Context {
	return w.ctx
}

// streamServerInterceptor is a gRPC stream server interceptor
func (s *Server) streamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx, cancel := ic.Merge(ss.Context(), s.baseCtx)
		defer cancel()
		md, _ := grpcmd.FromIncomingContext(ctx)
		replyHeader := grpcmd.MD{}
		ctx = transport.NewServerContext(ctx, &Transport{
			endpoint:    s.endpoint.String(),
			operation:   info.FullMethod,
			reqHeader:   headerCarrier(md),
			replyHeader: headerCarrier(replyHeader),
		})

		h := func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx = context.WithValue(ctx, stream{
				ServerStream: ss,
				middlware:    s.middleware,
			}, ss)
			return handler(srv, ss), nil
		}
		if next := s.middleware.Match(info.FullMethod); len(next) > 0 {
			h = middleware.Chain(next...)(h)
		}

		ws := NewWrappedStream(ctx, ss, s.middleware)

		err := handler(srv, ws)
		if len(replyHeader) > 0 {
			_ = grpc.SetHeader(ctx, replyHeader)
		}
		return err
	}
}

type stream struct {
	grpc.ServerStream
	middlware matcher.Matcher
}

func GetStream(ctx context.Context) grpc.ServerStream {
	return ctx.Value(stream{}).(grpc.ServerStream)
}

func (w *wrappedStream) SendMsg(m interface{}) error {
	err := w.ServerStream.SendMsg(m)
	info, _ := transport.FromServerContext(w.ctx)

	h := func(ctx context.Context, req interface{}) (interface{}, error) {
		return req, nil
	}

	if next := w.middleware.Match(info.Operation()); len(next) > 0 {
		h = middleware.Chain(next...)(h)
	}

	_, err = h(w.ctx, m)

	return err
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
	err := w.ServerStream.RecvMsg(m)
	info, _ := transport.FromServerContext(w.ctx)

	h := func(ctx context.Context, req interface{}) (interface{}, error) {
		return req, nil
	}

	if next := w.middleware.Match(info.Operation()); len(next) > 0 {
		h = middleware.Chain(next...)(h)
	}

	_, err = h(w.ctx, m)
	return err
}
