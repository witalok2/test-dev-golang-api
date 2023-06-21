package api

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/witalok2/test-dev-golang-api/adapter/primary/handler"
	"github.com/witalok2/test-dev-golang-api/config"
)

func NewRoute(ctx context.Context, e *echo.Echo, environment *config.Environment, handler *handler.Handler) {
	e.GET("/healthcheck", func(c echo.Context) error { return c.String(http.StatusOK, "online") })

	v1 := e.Group("/v1")
	Swagger(Options{Port: environment.API.Port, Group: v1.Group("/docs/*")})

	client := v1.Group("/client")
	client.GET("", handler.ListClient(ctx))
	client.GET("/:id", handler.GetClient(ctx))
	client.POST("", handler.CreateClient(ctx))
	client.PUT("/:id", handler.UpdateClient(ctx))
	client.DELETE("/:id", handler.DeleteClient(ctx))
}
