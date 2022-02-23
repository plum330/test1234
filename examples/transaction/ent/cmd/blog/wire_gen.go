// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/examples/transaction/ent/internal/biz"
	"github.com/go-kratos/kratos/examples/transaction/ent/internal/conf"
	"github.com/go-kratos/kratos/examples/transaction/ent/internal/data"
	"github.com/go-kratos/kratos/examples/transaction/ent/internal/server"
	"github.com/go-kratos/kratos/examples/transaction/ent/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	cardRepo := data.NewCardRepo(dataData, logger)
	transaction := data.NewTransaction(dataData)
	articleUsecase := biz.NewArticleUsecase(userRepo, cardRepo, transaction, logger)
	blogService := service.NewTransactionService(articleUsecase, logger)
	httpServer := server.NewHTTPServer(confServer, logger, blogService)
	grpcServer := server.NewGRPCServer(confServer, logger, blogService)
	app := newApp(logger, httpServer, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
