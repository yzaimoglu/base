package view

import (
	"context"
	"net/url"
	"os"
	"path"
	"strconv"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v5"
	"github.com/rs/zerolog/log"
	"github.com/yzaimoglu/base/utils/crypto"
)

type View struct {
}

func NewView() *View {
	return &View{}
}

func (v *View) Handle(c echo.Context, cmp templ.Component) error {
	v.saveLocallyDev(cmp, c.Request().URL)
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func (v *View) saveLocallyDev(cmp templ.Component, url *url.URL) error {
	debug := os.Getenv("DEBUG")
	debugBool, err := strconv.ParseBool(debug)
	if err != nil {
		debugBool = false
	}

	if debugBool {
		rootPath := "./ui/html/"

		dir := path.Join(rootPath)
		if err := os.MkdirAll(dir, 0755); err != nil && err != os.ErrExist {
			log.Fatal().Msgf("failed to create dir %q: %v", dir, err)
		}
		fileName := crypto.Sha256(url.Path) + ".html"

		filePath := path.Join(dir, fileName)
		f, err := os.Create(filePath)
		if err != nil {
			log.Fatal().Msgf("failed to create output file: %v", err)
		}

		return cmp.Render(context.Background(), f)
	}
	return nil
}
