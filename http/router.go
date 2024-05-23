package http

import (
	net_http "net/http"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"github.com/rs/zerolog/log"
	"github.com/yzaimoglu/base/base"
	"github.com/yzaimoglu/base/http/view"
	config_utils "github.com/yzaimoglu/base/utils/config"
	http_utils "github.com/yzaimoglu/base/utils/http"
)

type Router struct {
	Base       *base.Base
	Echo       *echo.Echo
	View       *view.View
	Controller *Controller
}

func NewRouter(e *echo.Echo, b *base.Base) *Router {
	return &Router{
		Base:       b,
		Echo:       e,
		View:       view.NewView(),
		Controller: NewController(),
	}
}

func (r *Router) Setup() {
	r.MiddlewareSetup()
	r.Echo.GET("/api/v1/test", r.Controller.API.Test)
	r.Echo.GET("/", r.Controller.Home.Index)
}

func (r *Router) MiddlewareSetup() {
	r.Echo.Pre(middleware.RemoveTrailingSlash())
	log.Info().Msg("Enabling custom context middleware.")
	r.Echo.Use(http_utils.EnableCustomContext)
	log.Info().Msg("Enabling request id middleware.")
	r.Echo.Use(middleware.RequestID())
	log.Info().Msg("Enabling recover middleware.")
	r.Echo.Use(middleware.Recover())
	log.Info().Msg("Enabling secure middleware.")
	r.Echo.Use(middleware.Secure())
	log.Info().Msg("Enabling csrf middleware.")
	r.Echo.Use(middleware.CSRF())
	limitMB := 50
	log.Info().Msgf("Enabling body limit middleware set to %d MB.", limitMB)
	r.Echo.Use(middleware.BodyLimit(50 * 1024 * 1024)) // 50 MB Body Limit
	log.Info().Msg("Enabling context value middleware.")
	r.Echo.Use(http_utils.SetContextValues)
	log.Info().Msg("Enabling logger middleware.")
	r.Echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Info().
				Str("type", "http").
				Int("status", v.Status).
				Str("uri", v.URI).
				Str("request_id", c.Get(http_utils.KeyRequestID).(string)).
				Str("csrf_token", c.Get(http_utils.KeyCSRF).(string)).
				Msgf("received.")
			return nil
		},
	}))

	log.Info().Msg("Registering static assets.")
	if config_utils.IsDebug() {
		log.Debug().Msg("Disabling cache for static assets.")
		r.Echo.GET("/static/*", echo.WrapHandler(net_http.StripPrefix("/static", net_http.FileServer(net_http.Dir("./ui/static")))), http_utils.NoCache, middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}))
	} else {
		r.Echo.GET("/static/*", echo.WrapHandler(net_http.StripPrefix("/static", net_http.FileServer(net_http.Dir("./ui/static")))), middleware.GzipWithConfig(middleware.GzipConfig{
			Level: 5,
		}))
	}

}
