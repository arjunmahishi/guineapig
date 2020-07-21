package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

var (
	delay          = flag.Duration("delay", 0, "delay string. EG: '1s', '1ms' etc")
	port           = flag.String("port", "3000", "port")
	response       = flag.String("response", "", "response file")
	count          = 0
	responseString json.RawMessage
)

func main() {
	flag.Parse()

	readResponse()

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
	log.Println(c.Request())
	time.Sleep(*delay)
	return c.JSON(http.StatusOK, responseString)
}

func readResponse() {
	if *response != "" {
		raw, err := ioutil.ReadFile(*response)
		if err != nil {
			panic(err)
		}
		responseString = raw
	}
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
