package controller

import (
	"github.com/labstack/echo/v5"
	"github.com/yzaimoglu/base/types"
)

type ApiController struct {
}

func NewApiController() *ApiController {
	return &ApiController{}
}

func (c *ApiController) Test(ctx echo.Context) error {
	return ctx.JSON(200, types.JSON{
		"test": "test",
	})
}
