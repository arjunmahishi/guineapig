package core

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

func handler(c echo.Context) error {
	pathMap, ok := Config[c.Path()]
	if !ok {
		return c.JSON(http.StatusNotFound, json.RawMessage("{}"))
	}
	response, ok := pathMap[c.Request().Method]
	if !ok {
		return c.JSON(http.StatusMethodNotAllowed, json.RawMessage("{}"))
	}
	return c.JSON(response.StatusCode, response.Response)
}

// TODO: support different content-types
