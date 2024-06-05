package controller

import (
	"github.com/labstack/echo/v5"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
	"github.com/yzaimoglu/base/base"
	"github.com/yzaimoglu/base/http/view"
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
	// title := "404: Not Found"
	// description := "The page you were looking for could not be found."
	// alpine := false
	// htmx := false
	return c.View.NotFound(ctx, html.Div(gomponents.Text("Not Found")))
}

func (c *ErrorController) BadRequest(ctx echo.Context) error {
	// title := "400: Bad Request"
	// description := "Your request could not be sent."
	// alpine := false
	// htmx := false
	return c.View.BadRequest(ctx, html.Div(gomponents.Text("Bad Request")))
}

func (c *ErrorController) InternalServerError(ctx echo.Context) error {
	// title := "500: Internal Server error"
	// description := "An internal server occured, please report to the site administrator."
	// alpine := false
	// htmx := false
	return c.View.InternalServerError(ctx, html.Div(gomponents.Text("Internal Server Error")))
}
