package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

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
	e.GET("*", handle)
	e.HideBanner = true
	printBanner()
	e.Logger.Fatal(e.Start(":" + *port))
}

func handle(c echo.Context) error {
	count++
	var b json.RawMessage
	switch c.Request().Method {
	case http.MethodPost:
		if err := c.Bind(&b); err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("Body: ", string(b), "request-count:", count, "path:", c.Path())
	time.Sleep(*delay)
	return c.String(http.StatusOK, "OK")
}

func printBanner() {
	print(string("\033[35m"), `
 ██████╗ ██╗   ██╗██╗███╗   ██╗███████╗ █████╗ ██████╗ ██╗ ██████╗ 
██╔════╝ ██║   ██║██║████╗  ██║██╔════╝██╔══██╗██╔══██╗██║██╔════╝ 
██║  ███╗██║   ██║██║██╔██╗ ██║█████╗  ███████║██████╔╝██║██║  ███╗
██║   ██║██║   ██║██║██║╚██╗██║██╔══╝  ██╔══██║██╔═══╝ ██║██║   ██║
╚██████╔╝╚██████╔╝██║██║ ╚████║███████╗██║  ██║██║     ██║╚██████╔╝
 ╚═════╝  ╚═════╝ ╚═╝╚═╝  ╚═══╝╚══════╝╚═╝  ╚═╝╚═╝     ╚═╝ ╚═════╝ 
`, string("\033[0m"))
}
