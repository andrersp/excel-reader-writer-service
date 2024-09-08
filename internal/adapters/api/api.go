package api

import (
	"context"
	"escrituras/internal/adapters/api/middlewares"
	"escrituras/internal/adapters/config"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type api struct {
	server *echo.Echo
}

func (a *api) Add(method, path string, handler echo.HandlerFunc, middlewares ...echo.MiddlewareFunc) {
	a.server.Add(method, path, handler, middlewares...)
}

func (a *api) Start() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go func() {
		if err := a.server.Start(fmt.Sprintf(":%s", config.APP_PORT)); err != nil && err != http.ErrServerClosed {
			a.server.Logger.Fatal("shutting down ther server")
		}
	}()
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return a.server.Shutdown(ctx)

}

func NewApi() *api {
	e := echo.New()
	e.Use(middleware.Recover())
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Use(middlewares.LoggerMiddleware)
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	return &api{server: e}
}
