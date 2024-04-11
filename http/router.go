package http

import (
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/yzaimoglu/base/base"
	"github.com/yzaimoglu/base/types"
	"github.com/yzaimoglu/base/ui"
)

type Router struct {
	Base *base.Base
	Echo *echo.Echo
}

func NewRouter(e *echo.Echo, b *base.Base) *Router {
	return &Router{
		Base: b,
		Echo: e,
	}
}

func (r *Router) Setup() {
	r.Echo.GET("/api/v1/test", r.GetTestRoute)

	r.setupFrontend()
}

func (r *Router) setupFrontend() {
	r.Echo.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: ui.BuildHTTPFS(),
		Skipper: func(c echo.Context) bool {
			// Exclude Pocketbase paths from being overwritten by Svelte paths
			if strings.HasPrefix(c.Path(), "/api/") || strings.HasPrefix(c.Path(), "/_/") {
				return true
			}
			return false
		},
		HTML5: true,
	}))
}

func (r *Router) GetTestRoute(c echo.Context) error {
	return c.JSON(200, types.JSON{
		"test": "test",
	})
}
