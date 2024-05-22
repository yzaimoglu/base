package config

import (
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase/core"
	"github.com/rs/zerolog/log"
	"github.com/yzaimoglu/base/base"
	"github.com/yzaimoglu/base/http"
)

func ConfigurePocketbase(a *base.Base) {
	log.Info().Msg("Configuring pocketbase...")
	a.Pocketbase.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		return ConfigureRouter(e.Router, a)
	})
	log.Info().Msg("Pocketbase configured.")
}

func ConfigureRouter(e *echo.Echo, a *base.Base) error {
	router := http.NewRouter(e, a)
	router.Echo.Static("/static", "./ui/static")
	router.Setup()
	return nil
}
