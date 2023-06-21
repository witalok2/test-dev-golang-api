package api

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	logger "github.com/sirupsen/logrus"
	"github.com/witalok2/test-dev-golang-api/internal/entity"
)

var echoServer *echo.Echo

func New() *echo.Echo {
	echoServer = echo.New()

	if os.Getenv("GAE_ENV") == entity.PRODUCTION {
		echoServer.Use(middleware.CORS())      // Enable CORS on API
		echoServer.Use(middleware.Recover())   // Recover the API on fatal situations
		echoServer.Use(middleware.RequestID()) // Add a unique ID in each request
		echoServer.Use(middleware.Gzip())      // Compress the request with Gzip

		logger.Info("cors enabled")
	} else {
		echoServer.Use(middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"},
			AllowHeaders: []string{"*"},
			AllowMethods: []string{"*"},
		}))
	}

	return echoServer
}

func ProvideEchoInstance(task func(e *echo.Echo)) {
	task(echoServer)
}

func Run(port string) {
	echoServer.Logger.Fatal(echoServer.Start(":" + port))
}

func Stop() {
	echoServer.Close()
}

func Use(middleware ...echo.MiddlewareFunc) {
	echoServer.Use(middleware...)
}
