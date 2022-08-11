package main

import (
	"fmt"
	"net/url"
	"os"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	target  = os.Getenv("TARGET_URL")
	Version = "dev"
)

func main() {
	fmt.Printf("http-proxy-metrics-collector version: %s\n", Version)
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Setup proxy
	url1, err := url.Parse(target)
	if err != nil {
		e.Logger.Fatal(err)
	}
	targets := []*middleware.ProxyTarget{
		{
			URL: url1,
		},
	}

	// Setup prometheus collector
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)

	// Setup RR
	e.Group("/").Use(middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

	e.Logger.Fatal(e.Start("0.0.0.0:3000"))
}
