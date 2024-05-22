package http

import (
	"github.com/yzaimoglu/base/http/controller"
	"github.com/yzaimoglu/base/http/view"
)

type Controller struct {
	Home *controller.HomeController
	API  *controller.ApiController
}

func NewController() *Controller {
	v := view.NewView()
	return &Controller{
		Home: controller.NewHomeController(v),
		API:  controller.NewApiController(),
	}
}
