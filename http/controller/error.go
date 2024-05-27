package controller

import (
	"github.com/labstack/echo/v5"
	"github.com/yzaimoglu/base/base"
	"github.com/yzaimoglu/base/http/view"
	"github.com/yzaimoglu/base/ui/layout"

	error_ui "github.com/yzaimoglu/base/ui/error"
)

type ErrorController struct {
	Base *base.Base
	View *view.View
}

func NewErrorController(b *base.Base, v *view.View) *ErrorController {
	return &ErrorController{
		Base: b,
		View: v,
	}
}

func (c *ErrorController) NotFound(ctx echo.Context) error {
	title := "404: Not Found"
	description := "The page you were looking for could not be found."
	alpine := false
	htmx := false
	return c.View.Handle(ctx, layout.Base(title, description, alpine, htmx, error_ui.NotFound()))
}

func (c *ErrorController) BadRequest(ctx echo.Context) error {
	title := "400: Bad Request"
	description := "Your request could not be sent."
	alpine := false
	htmx := false
	return c.View.Handle(ctx, layout.Base(title, description, alpine, htmx, error_ui.BadRequest()))
}

func (c *ErrorController) InternalServerError(ctx echo.Context) error {
	title := "500: Internal Server error"
	description := "An internal server occured, please report to the site administrator."
	alpine := false
	htmx := false
	return c.View.Handle(ctx, layout.Base(title, description, alpine, htmx, error_ui.InternalServerError()))
}
