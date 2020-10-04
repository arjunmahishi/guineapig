package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func initRoutes(e *echo.Echo) {
	for path, pathMap := range config {
		for method := range pathMap {
			createRoute(e, path, method)
		}
	}
}

func createRoute(e *echo.Echo, path, method string) {
	switch method {
	case http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete:
		e.GET(path, handler)
	default:
		log.Fatalf(
			"the method '%s' is not supported (yet). If it is a valid method, make a feature request (%s)",
			method,
			"https://github.com/arjunmahishi/guineapig/issues/new",
		)
	}
}
