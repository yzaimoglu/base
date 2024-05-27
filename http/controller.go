package http

import (
	"github.com/yzaimoglu/base/base"
	"github.com/yzaimoglu/base/http/controller"
	"github.com/yzaimoglu/base/http/view"
)

type Controller struct {
	Base  *base.Base
	Home  *controller.HomeController
	Error *controller.ErrorController
	API   *controller.ApiController
}

func NewController(b *base.Base) *Controller {
	v := view.NewView()
	return &Controller{
		Base:  b,
		Home:  controller.NewHomeController(b, v),
		Error: controller.NewErrorController(b, v),
		API:   controller.NewApiController(b),
	}
}
