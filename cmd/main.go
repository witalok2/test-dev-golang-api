package main

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
	"github.com/witalok2/test-dev-golang-api/adapter/primary/api"
	"github.com/witalok2/test-dev-golang-api/adapter/primary/handler"
	"github.com/witalok2/test-dev-golang-api/adapter/secondary/database"
	amqp "github.com/witalok2/test-dev-golang-api/adapter/secondary/rabbitMQ"
	"github.com/witalok2/test-dev-golang-api/config"
	"github.com/witalok2/test-dev-golang-api/internal/entity"
	"github.com/witalok2/test-dev-golang-api/internal/service"
)

func init() {
	logger.New().WithContext(context.WithValue(context.Background(), entity.SERVICE, entity.SERVICE_NAME))
}

func main() {
	ctx := context.Background()
	logger.Infof("Starting execution: %v", time.Now().Format("2006-01-02 15:04:05"))

	env, err := config.LoadConfig()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("instantiating database")
	dbReader, err := database.NewReaderConnection(env.Database.URI)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("instantiating message queue")
	rabbitMQ, err := amqp.NewQueueClient(env.Queue)
	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("instantiating service")
	service := service.NewService(ctx, dbReader, rabbitMQ)
	handler := handler.NewHandler(service, env)

	logger.Info("instantiating api")
	api.New()
	api.ProvideEchoInstance(func(e *echo.Echo) { api.NewRoute(ctx, e, env, handler) })
	api.Run(env.API.Port)
}
