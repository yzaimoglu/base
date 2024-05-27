package http_utils

import (
	"errors"
	"fmt"
	"net/http"

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

func NewHTTPErrorHandler(exposeError bool) echo.HTTPErrorHandler {
	return func(c echo.Context, err error) {
		if c.Response().Committed {
			return
		}

		he := &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusText(http.StatusInternalServerError),
		}
		if errors.As(err, &he) {
			if he.Internal != nil { // max 2 levels of checks even if internal could have also internal
				errors.As(he.Internal, &he)
			}
		}

		// Issue #1426
		code := he.Code
		// message := he.Message
		// switch m := he.Message.(type) {
		// case string:
		// 	if exposeError {
		// 		message = echo.Map{"message": m, "error": err.Error()}
		// 	} else {
		// 		message = echo.Map{"message": m}
		// 	}
		// case json.Marshaler:
		// 	// do nothing - this type knows how to format itself to JSON
		// case error:
		// 	message = echo.Map{"message": m.Error()}
		// }

		// Send response
		var cErr error
		if c.Request().Method == http.MethodHead { // Issue #608
			// cErr = c.NoContent(he.Code)
			cErr = c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("/%d", code))
		} else {
			// cErr = c.JSON(code, message)
			cErr = c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf("/%d", code))
		}
		if cErr != nil {
			c.Echo().Logger.Error(err) // truly rare case. ala client already disconnected
		}
	}
}
