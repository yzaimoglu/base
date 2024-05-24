package controller

import (
	"github.com/labstack/echo/v5"
	"github.com/yzaimoglu/base/base"
	"github.com/yzaimoglu/base/types"
)

type ApiController struct {
	Base *base.Base
}

func NewApiController(b *base.Base) *ApiController {
	return &ApiController{
		Base: b,
	}
}

func (c *ApiController) Test(ctx echo.Context) error {
	return ctx.JSON(200, types.JSON{
		"test": "test",
	})
}
