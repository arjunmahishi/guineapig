package main

import (
	"flag"
	"log"

	"github.com/labstack/echo"
)

var (
	port     = flag.String("port", "3000", "port")
	response = flag.String("response", "", "response file")
)

func main() {
	flag.Parse()
	if err := initConfig(); err != nil {
		log.Fatalln(err.Error())
	}

	e := echo.New()
	initRoutes(e, config)
	e.HideBanner = true
	printBanner()
	e.Logger.Fatal(e.Start(":" + *port))
}

func initRoutes(e *echo.Echo, routes []route) {
	for _, r := range routes {
		e.Add(r.Method, r.Path, func(c echo.Context) error {
			req := c.Request()
			log.Println(req.URL, r.ResponseStatusCode)
			return c.JSON(r.ResponseStatusCode, r.ResponseBody)
		})
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
