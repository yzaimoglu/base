package controller

import (
	"github.com/labstack/echo/v5"
	"github.com/yzaimoglu/base/http/view"
	"github.com/yzaimoglu/base/ui/home"
	"github.com/yzaimoglu/base/ui/layout"
)

type HomeController struct {
	View *view.View
}

func NewHomeController(v *view.View) *HomeController {
	return &HomeController{
		View: v,
	}
}

func (c *HomeController) Show(ctx echo.Context) error {
	title := "Test"
	description := "Test Description"
	alpine := true
	return c.View.Handle(ctx, layout.Base(title, description, alpine, home.Show()))
}

func (c *HomeController) Index(ctx echo.Context) error {
	title := "Test"
	description := "Test Description"
	alpine := true
	return c.View.Handle(ctx, layout.Base(title, description, alpine, home.Show()))
}
