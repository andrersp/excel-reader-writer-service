package middlewares

import (
	"log/slog"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		start := time.Now()
		path := c.Request().RequestURI

		method := c.Request().Method
		if err := next(c); err != nil {
			c.Error(err)
		}
		if strings.Contains(path, "/docs/") {
			return next(c)
		}
		end := time.Since(start)

		status := c.Response().Status
		keys := []interface{}{
			"duration", end.Seconds(), "url", path, "method", method, "status", status,
		}
		if status <= 400 {
			slog.Info("sucess", keys...)
		} else {
			slog.Error("error", keys...)
		}

		return nil
	}
}
