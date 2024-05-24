package controller

import (
	"github.com/labstack/echo/v5"
	"github.com/yzaimoglu/base/base"
	"github.com/yzaimoglu/base/http/view"
	"github.com/yzaimoglu/base/ui/home"
	"github.com/yzaimoglu/base/ui/layout"
)

type HomeController struct {
	Base *base.Base
	View *view.View
}

func NewHomeController(b *base.Base, v *view.View) *HomeController {
	return &HomeController{
		Base: b,
		View: v,
	}
}

func (c *HomeController) Show(ctx echo.Context) error {
	title := "Test"
	description := "Test Description"
	alpine := true
	htmx := true
	return c.View.Handle(ctx, layout.Base(title, description, alpine, htmx, home.Show()))
}

func (c *HomeController) Test(ctx echo.Context) error {
	title := "Test"
	description := "Test Description"
	alpine := true
	htmx := true
	return c.View.Handle(ctx, layout.Base(title, description, alpine, htmx, home.Test()))
}

func (c *HomeController) Index(ctx echo.Context) error {
	title := "Test"
	description := "Test Description"
	alpine := true
	htmx := true
	return c.View.Handle(ctx, layout.Base(title, description, alpine, htmx, home.Show()))
}
