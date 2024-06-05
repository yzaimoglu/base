package controller

import (
	"github.com/labstack/echo/v5"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
	"github.com/yzaimoglu/base/base"
	"github.com/yzaimoglu/base/http/view"
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
	return c.View.Success(ctx, c.View.Page(
		ctx,
		"Index",
		"Index Page",
		html.Div(components.Classes(map[string]bool{
			"bg-red-800":     true,
			"text-sm":        false,
			"justify-center": true,
			"flex":           true,
			"items-center":   true,
		}), gomponents.Text("Test Page")),
	))
}

func (c *HomeController) Franken(ctx echo.Context) error {
	return c.View.NotFound(ctx, c.View.Page(
		ctx,
		"Franken",
		"Franken Page",
		html.Div(gomponents.Text("Franken Page")),
	))
}
