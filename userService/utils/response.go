package utils

import (
	"github.com/labstack/echo/v4"
)

func Response(c echo.Context, code int, name string, data any) error {
	return c.JSON(code, map[string]interface{}{
		"result": true,
		name:     data,
	})
}

func ErrorResponse(c echo.Context, code int, err error) error {
	return c.JSON(code, map[string]interface{}{"error": []string{err.Error()}})
}
