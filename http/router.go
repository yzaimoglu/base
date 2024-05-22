package http

import (
	"github.com/labstack/echo/v5"
	"github.com/yzaimoglu/base/base"
	"github.com/yzaimoglu/base/http/view"
)

type Router struct {
	Base       *base.Base
	Echo       *echo.Echo
	View       *view.View
	Controller *Controller
}

func NewRouter(e *echo.Echo, b *base.Base) *Router {
	return &Router{
		Base:       b,
		Echo:       e,
		View:       view.NewView(),
		Controller: NewController(),
	}
}

func (r *Router) Setup() {
	r.Echo.GET("/api/v1/test", r.Controller.API.Test)
	r.Echo.GET("/", r.Controller.Home.Index)
}
