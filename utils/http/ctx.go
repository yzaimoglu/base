package http_utils

import (
	"context"

	"github.com/labstack/echo/v5"
)

type CustomEchoContext struct {
	echo.Context
}

func (c CustomEchoContext) Get(key string) interface{} {
	// grab value from echo.Context
	val := c.Context.Get(key)
	// if it's not nil, return it
	if val != nil {
		return val
	}
	// otherwise, return Request.Context
	return c.Request().Context().Value(key)
}

func (c CustomEchoContext) Set(key string, val interface{}) {
	// we're replacing the whole Request in echo.Context
	// with a copied request that has the updated context value
	c.SetRequest(
		c.Request().WithContext(
			context.WithValue(c.Request().Context(), key, val),
		),
	)
}

func GetString(ctx context.Context, key string) string {
	if v, ok := ctx.Value(key).(string); ok {
		return v
	}
	return ""
}

const (
	KeyCSRF        = "csrf"
	KeyRequestID   = "request_id"
	KeyCurrentPath = "current_path"
	KeyUserAgent   = "user_agent"
)

func CSRF(ctx context.Context) string {
	return GetString(ctx, KeyCSRF)
}

func RequestID(ctx context.Context) string {
	return GetString(ctx, KeyRequestID)
}

func CurrentPath(ctx context.Context) string {
	return GetString(ctx, KeyCurrentPath)
}

func UserAgent(ctx context.Context) string {
	return GetString(ctx, KeyUserAgent)
}
