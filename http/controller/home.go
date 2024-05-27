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

func (c *HomeController) Index(ctx echo.Context) error {
	title := "Base"
	description := "A starter project for the GOTH Stack with Franken UI"
	alpine := true
	htmx := true
	return c.View.Handle(ctx, layout.Base(title, description, alpine, htmx, home.Index()))
}

func (c *HomeController) Franken(ctx echo.Context) error {
	title := "Franken UI"
	description := "Implementation of Franken UI components in Go Templ"
	alpine := false
	htmx := true
	return c.View.Handle(ctx, layout.Base(title, description, alpine, htmx, home.Franken()))
}
