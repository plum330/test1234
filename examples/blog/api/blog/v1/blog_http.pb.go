// Code generated by protoc-gen-go-http. DO NOT EDIT.

package v1

import (
	context "context"
	middleware "github.com/go-kratos/kratos/v2/middleware"
	http1 "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	mux "github.com/gorilla/mux"
	http "net/http"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(http.Request)
var _ = new(context.Context)
var _ = new(middleware.Middleware)
var _ = binding.BindVars
var _ = mux.NewRouter

const _ = http1.SupportPackageIsVersion1

type BlogServiceHandler interface {
	CreateArticle(context.Context, *CreateArticleRequest) (*CreateArticleReply, error)

	DeleteArticle(context.Context, *DeleteArticleRequest) (*DeleteArticleReply, error)

	GetArticle(context.Context, *GetArticleRequest) (*GetArticleReply, error)

	ListArticle(context.Context, *ListArticleRequest) (*ListArticleReply, error)

	UpdateArticle(context.Context, *UpdateArticleRequest) (*UpdateArticleReply, error)
}

func NewBlogServiceHandler(srv BlogServiceHandler, opts ...http1.HandleOption) http.Handler {
	h := http1.DefaultHandleOptions()
	for _, o := range opts {
		o(&h)
	}
	r := mux.NewRouter()

	r.HandleFunc("/v1/article/", func(w http.ResponseWriter, r *http.Request) {
		var in CreateArticleRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateArticle(ctx, req.(*CreateArticleRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		ctx := middleware.WithMethod(r.Context(), "/blog.api.v1.BlogService/CreateArticle")
		out, err := next(ctx, &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*CreateArticleReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("POST")

	r.HandleFunc("/v1/article/{id}", func(w http.ResponseWriter, r *http.Request) {
		var in UpdateArticleRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.BindVars(mux.Vars(r), &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateArticle(ctx, req.(*UpdateArticleRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		ctx := middleware.WithMethod(r.Context(), "/blog.api.v1.BlogService/UpdateArticle")
		out, err := next(ctx, &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*UpdateArticleReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("PUT")

	r.HandleFunc("/v1/article/{id}", func(w http.ResponseWriter, r *http.Request) {
		var in DeleteArticleRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.BindVars(mux.Vars(r), &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteArticle(ctx, req.(*DeleteArticleRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		ctx := middleware.WithMethod(r.Context(), "/blog.api.v1.BlogService/DeleteArticle")
		out, err := next(ctx, &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*DeleteArticleReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("DELETE")

	r.HandleFunc("/v1/article/{id}", func(w http.ResponseWriter, r *http.Request) {
		var in GetArticleRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		if err := binding.BindVars(mux.Vars(r), &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetArticle(ctx, req.(*GetArticleRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		ctx := middleware.WithMethod(r.Context(), "/blog.api.v1.BlogService/GetArticle")
		out, err := next(ctx, &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*GetArticleReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	r.HandleFunc("/v1/article/", func(w http.ResponseWriter, r *http.Request) {
		var in ListArticleRequest
		if err := h.Decode(r, &in); err != nil {
			h.Error(w, r, err)
			return
		}

		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListArticle(ctx, req.(*ListArticleRequest))
		}
		if h.Middleware != nil {
			next = h.Middleware(next)
		}
		ctx := middleware.WithMethod(r.Context(), "/blog.api.v1.BlogService/ListArticle")
		out, err := next(ctx, &in)
		if err != nil {
			h.Error(w, r, err)
			return
		}
		reply := out.(*ListArticleReply)
		if err := h.Encode(w, r, reply); err != nil {
			h.Error(w, r, err)
		}
	}).Methods("GET")

	return r
}

type BlogServiceHTTPClient interface {
	CreateArticle(ctx context.Context, req *CreateArticleRequest, opts ...http1.CallOption) (rsp *CreateArticleReply, err error)

	DeleteArticle(ctx context.Context, req *DeleteArticleRequest, opts ...http1.CallOption) (rsp *DeleteArticleReply, err error)

	GetArticle(ctx context.Context, req *GetArticleRequest, opts ...http1.CallOption) (rsp *GetArticleReply, err error)

	ListArticle(ctx context.Context, req *ListArticleRequest, opts ...http1.CallOption) (rsp *ListArticleReply, err error)

	UpdateArticle(ctx context.Context, req *UpdateArticleRequest, opts ...http1.CallOption) (rsp *UpdateArticleReply, err error)
}

type BlogServiceHTTPClientImpl struct {
	cc *http1.Client
}

func NewBlogServiceHTTPClient(client *http1.Client) BlogServiceHTTPClient {
	return &BlogServiceHTTPClientImpl{client}
}

func (c *BlogServiceHTTPClientImpl) CreateArticle(ctx context.Context, in *CreateArticleRequest, opts ...http1.CallOption) (*CreateArticleReply, error) {
	var out CreateArticleReply
	path := binding.EncodePath("POST", "/v1/article/", in)
	ctx = middleware.WithMethod(ctx, "/blog.api.v1.BlogService/CreateArticle")

	err := c.cc.Invoke(ctx, "POST", path, in, &out)

	return &out, err
}

func (c *BlogServiceHTTPClientImpl) DeleteArticle(ctx context.Context, in *DeleteArticleRequest, opts ...http1.CallOption) (*DeleteArticleReply, error) {
	var out DeleteArticleReply
	path := binding.EncodePath("DELETE", "/v1/article/{id}", in)
	ctx = middleware.WithMethod(ctx, "/blog.api.v1.BlogService/DeleteArticle")

	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out)

	return &out, err
}

func (c *BlogServiceHTTPClientImpl) GetArticle(ctx context.Context, in *GetArticleRequest, opts ...http1.CallOption) (*GetArticleReply, error) {
	var out GetArticleReply
	path := binding.EncodePath("GET", "/v1/article/{id}", in)
	ctx = middleware.WithMethod(ctx, "/blog.api.v1.BlogService/GetArticle")

	err := c.cc.Invoke(ctx, "GET", path, nil, &out)

	return &out, err
}

func (c *BlogServiceHTTPClientImpl) ListArticle(ctx context.Context, in *ListArticleRequest, opts ...http1.CallOption) (*ListArticleReply, error) {
	var out ListArticleReply
	path := binding.EncodePath("GET", "/v1/article/", in)
	ctx = middleware.WithMethod(ctx, "/blog.api.v1.BlogService/ListArticle")

	err := c.cc.Invoke(ctx, "GET", path, nil, &out)

	return &out, err
}

func (c *BlogServiceHTTPClientImpl) UpdateArticle(ctx context.Context, in *UpdateArticleRequest, opts ...http1.CallOption) (*UpdateArticleReply, error) {
	var out UpdateArticleReply
	path := binding.EncodePath("PUT", "/v1/article/{id}", in)
	ctx = middleware.WithMethod(ctx, "/blog.api.v1.BlogService/UpdateArticle")

	err := c.cc.Invoke(ctx, "PUT", path, in, &out)

	return &out, err
}
