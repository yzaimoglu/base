package view

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/components"
	"github.com/maragudk/gomponents/html"
)

type View struct {
}

func NewView() *View {
	return &View{}
}

func (v *View) Page(ctx echo.Context, title, description string, body gomponents.Node) gomponents.Node {
	return components.HTML5(components.HTML5Props{
		Title:       title,
		Description: description,
		Head: []gomponents.Node{
			html.Meta(html.Charset("utf-8")),
			html.Meta(html.Name("viewport"), html.Content("width=device-width, initial-scale=1.0")),
			html.Link(html.Rel("apple-touch-icon"), gomponents.Attr("sizes", "60x60"), html.Href("/static/apple-icon-60x60.png")),
			html.Link(html.Rel("apple-touch-icon"), gomponents.Attr("sizes", "72x72"), html.Href("/static/apple-icon-72x72.png")),
			html.Link(html.Rel("apple-touch-icon"), gomponents.Attr("sizes", "76x76"), html.Href("/static/apple-icon-76x76.png")),
			html.Link(html.Rel("apple-touch-icon"), gomponents.Attr("sizes", "114x114"), html.Href("/static/apple-icon-114x114.png")),
			html.Link(html.Rel("apple-touch-icon"), gomponents.Attr("sizes", "120x120"), html.Href("/static/apple-icon-120x120.png")),
			html.Link(html.Rel("apple-touch-icon"), gomponents.Attr("sizes", "144x144"), html.Href("/static/apple-icon-144x144.png")),
			html.Link(html.Rel("apple-touch-icon"), gomponents.Attr("sizes", "152x152"), html.Href("/static/apple-icon-152x152.png")),
			html.Link(html.Rel("apple-touch-icon"), gomponents.Attr("sizes", "180x180"), html.Href("/static/apple-icon-180x180.png")),
			html.Link(html.Rel("icon"), html.Type("image/png"), gomponents.Attr("sizes", "192x192"), html.Href("/static/android-icon-192x192.png")),
			html.Link(html.Rel("icon"), html.Type("image/png"), gomponents.Attr("sizes", "32x32"), html.Href("/static/favicon-32x32.png")),
			html.Link(html.Rel("icon"), html.Type("image/png"), gomponents.Attr("sizes", "96x96"), html.Href("/static/favicon-96x96.png")),
			html.Link(html.Rel("icon"), html.Type("image/png"), gomponents.Attr("sizes", "16x16"), html.Href("/static/favicon-16x16.png")),
			html.Meta(html.Name("msapplication-TileColor"), html.Content("#ffffff")),
			html.Meta(html.Name("msapplication-TileImage"), html.Content("/static/ms-icon-144x144.png")),
			html.Meta(html.Name("theme-color"), html.Content("#ffffff")),
			gomponents.If(true, html.Script(html.Async(), html.Src("https://cdn.jsdelivr.net/npm/uikit@3.20.8/dist/js/uikit.min.js"))),
			gomponents.If(true, html.Script(html.Async(), html.Src("https://cdn.jsdelivr.net/npm/uikit@3.20.8/dist/js/uikit-icons.min.js"))),
			html.Link(html.Rel("stylesheet"), html.Href("/static/main.css")),
			html.Link(html.Rel("stylesheet"), html.Href("https://unpkg.com/tailwindcss@2.1.2/dist/components.min.css")),
			html.Link(html.Rel("stylesheet"), html.Href("https://unpkg.com/@tailwindcss/typography@0.4.0/dist/typography.min.css")),
			html.Link(html.Rel("stylesheet"), html.Href("https://unpkg.com/tailwindcss@2.1.2/dist/utilities.min.css")),
			gomponents.If(true, html.Script(html.Src("https://unpkg.com/htmx.org@latest"))),
			gomponents.If(true, html.Script(html.Src("https://unpkg.com/htmx.org@1.9.12/dist/ext/loading-states.js"))),
			gomponents.If(true, html.Script(html.Src("https://unpkg.com/htmx.org@1.9.12/dist/ext/response-targets.js"))),
			gomponents.If(true, html.Script(html.Defer(), html.Src("https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js"))),
			gomponents.If(true, html.Script(html.Defer(), html.Src("https://cdn.jsdelivr.net/npm/@alpinejs/collapse@3.x.x/dist/cdn.min.js"))),
			gomponents.If(true, html.Script(html.Src("https://unpkg.com/htmx.org@latest/dist/ext/alpine-morph.js"))),
			gomponents.If(true, html.Script(html.Defer(), html.Src("https://unpkg.com/@alpinejs/morph@3.x.x/dist/cdn.min.js"))),
		},
		Body: []gomponents.Node{
			body,
		},
	})
}

func (v *View) Handle(status int, c echo.Context, page gomponents.Node) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	c.Response().WriteHeader(status)
	return page.Render(c.Response().Writer)
}

func (v *View) Success(c echo.Context, page gomponents.Node) error {
	return v.Handle(http.StatusOK, c, page)
}

func (v *View) NotFound(c echo.Context, page gomponents.Node) error {
	return v.Handle(http.StatusNotFound, c, page)
}

func (v *View) BadRequest(c echo.Context, page gomponents.Node) error {
	return v.Handle(http.StatusBadRequest, c, page)
}

func (v *View) InternalServerError(c echo.Context, page gomponents.Node) error {
	return v.Handle(http.StatusInternalServerError, c, page)
}
