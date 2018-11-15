package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func main() {
	// Hosts
	hosts := map[string]*Host{}

	//-----
	// default
	//-----
	def := echo.New()
	def.Use(middleware.Logger())
	def.Use(middleware.Recover())

	def.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")
	})

	//------
	// localhost
	//------
	localhost := echo.New()
	localhost.Use(middleware.Logger())
	localhost.Use(middleware.Recover())

	hosts["localhost"] = &Host{localhost}

	localhost.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, localhost!")
	})

	//---------
	// example.com
	//---------
	example_com := echo.New()
	example_com.Use(middleware.Logger())
	example_com.Use(middleware.Recover())

	hosts["example.com"] = &Host{example_com}

	example_com.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, example.com!")
	})

	// Server
	e := echo.New()
	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		host := hosts[strings.Split(req.Host, ":")[0]]

		if host == nil {
			host = &Host{def}
		}

		host.Echo.ServeHTTP(res, req)

		return
	})
	e.Logger.Fatal(e.Start(":4000"))
}
