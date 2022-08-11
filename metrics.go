package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// MetricsCollector middleware collects prometheus metrics
func MetricsCollector(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		return next(c)
	}
}

// Metrics handler show page with prometheus metrics
func Metrics(c echo.Context) error {
	return c.String(http.StatusOK, "some metrics")
}
