package http

import (
	"github.com/yzaimoglu/base/base"
	"github.com/yzaimoglu/base/http/controller"
	"github.com/yzaimoglu/base/http/view"
)

type Controller struct {
	Base *base.Base
	Home *controller.HomeController
	API  *controller.ApiController
}

func NewController(b *base.Base) *Controller {
	v := view.NewView()
	return &Controller{
		Base: b,
		Home: controller.NewHomeController(b, v),
		API:  controller.NewApiController(b),
	}
}
