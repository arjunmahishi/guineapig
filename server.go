package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/davecgh/go-spew/spew"

	"github.com/labstack/echo"
)

var (
	delay = flag.Duration("delay", 0, "delay string. EG: '1s', '1ms' etc")
	port  = flag.String("port", "3000", "port")
	count = 0
)

func main() {
	flag.Parse()
	e := echo.New()
	e.POST("*", handle)
	e.PUT("*", handle)
	e.GET("*", func(c echo.Context) error {
		count++
		time.Sleep(*delay)
		fmt.Println("count:", count, "path:", c.Path())
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":" + *port))
}

func handle(c echo.Context) error {
	var b json.RawMessage
	if c.Request().Method == http.MethodPost {
		if err := c.Bind(&b); err != nil {
			fmt.Println(err)
		}
	}
	count++
	time.Sleep(*delay)
	fmt.Println(string(b), "count:", count, "path:", spew.Sdump(c.Request().Body))
	return c.String(http.StatusOK, "Hello, World!")
}
