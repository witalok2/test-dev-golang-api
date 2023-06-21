package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	logger "github.com/sirupsen/logrus"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/witalok2/test-dev-golang-api/config/docs"
)

type Options struct {
	Group     *echo.Group
	AccessKey string
	Port      string
}

func Swagger(opts Options) {
	docs.SwaggerInfo.Title = "Swagger test-dev-golang-api"
	docs.SwaggerInfo.Description = "Swagger com as rotas e modelos de uso de parâmetros da API do test-dev-golang-api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("localhost:%v", opts.Port) // Host of application
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	opts.Group.GET("", echoSwagger.WrapHandler)

	logger.Info("Swagger is initializing...")
}
