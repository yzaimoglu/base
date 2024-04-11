package config

import (
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase/core"
	"github.com/rs/zerolog/log"
	"github.com/yzaimoglu/base/base"
	"github.com/yzaimoglu/base/types"
	"github.com/yzaimoglu/base/ui"
)

func ConfigurePocketbase(a *base.Base) {
	log.Info().Msg("Configuring pocketbase...")
	a.Pocketbase.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		return ConfigureRouter(e.Router)
	})
	log.Info().Msg("Pocketbase configured.")
}

func ConfigureRouter(e *echo.Echo) error {
	e.GET("/api/v1/test", func(c echo.Context) error {
		return c.JSON(200, types.JSON{
			"test": "test",
		})
	})

	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Filesystem: ui.BuildHTTPFS(),
		Skipper: func(c echo.Context) bool {
			// Exclude Pocketbase paths from being overwritten by Svelte paths
			if strings.HasPrefix(c.Path(), "/api/") || strings.HasPrefix(c.Path(), "/_/") {
				return true
			}
			return false
		},
		HTML5: true,
	}))
	return nil
}
