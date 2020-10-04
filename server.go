package main

import (
	"flag"
	"log"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
)

var (
	port       = flag.String("port", "3000", "port")
	configFile = flag.String("config", "", "config file")
)

func main() {
	flag.Parse()
	if err := initConfig(); err != nil {
		log.Fatalln(err.Error(), "run with '-help' to get help")
	}

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method} ${uri} ${error}\n",
	}))
	initRoutes(e)
	e.HideBanner = true
	printBanner()
	e.Logger.Fatal(e.Start(":" + *port))
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
