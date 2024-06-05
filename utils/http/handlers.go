package http_utils

import (
	"github.com/labstack/echo/v5"
)

func NoCache(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate, max-age=0")
		c.Response().Header().Set("Pragma", "no-cache")
		c.Response().Header().Set("Expires", "0")
		return next(c)
	}
}

func SetContextValues(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		csrfCookie, err := c.Request().Cookie(KeyCSRF)
		if err == nil {
			c.Set(KeyCSRF, csrfCookie.Value)
		}
		c.Set(KeyRequestID, c.Response().Header().Get(echo.HeaderXRequestID))
		c.Set(KeyCurrentPath, c.Request().URL.Path)
		c.Set(KeyUserAgent, c.Request().UserAgent())
		return next(c)
	}
}

func EnableCustomContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		customCtx := CustomEchoContext{c}
		return next(customCtx)
	}
}
