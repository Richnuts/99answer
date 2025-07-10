package utils

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Success(c echo.Context, resp *http.Response) error {
	body, _ := io.ReadAll(resp.Body)
	var result interface{}
	json.Unmarshal(body, &result)

	return c.JSON(resp.StatusCode, result)
}
