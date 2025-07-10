package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

const RequestIDKey = "RequestID"

func RequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := c.Request().Header.Get("X-Request-ID")
			if requestID == "" {
				requestID = uuid.NewString()
			}

			// Set in the response header
			c.Response().Header().Set("X-Request-ID", requestID)

			// Save in context so handlers can use it
			c.Set(RequestIDKey, requestID)

			return next(c)
		}
	}
}
