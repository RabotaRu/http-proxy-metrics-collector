package main

import (
	"net/url"
	"os"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

// Version is http-proxy-metrics-collector binary version.
//nolint:gochecknoglobals
var Version = "dev"

func main() {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.INFO)
	e.Logger.Infof("http-proxy-metrics-collector version: %s", Version)

	// Setup proxy
	url1, err := url.Parse(os.Getenv("TARGET_URL"))
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
