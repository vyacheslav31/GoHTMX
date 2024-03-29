package app

import (
	"embed"
	"vyacheslav31/gohtmx/internal/templates/layouts"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type App struct {
	Environment string
}

func NewApp(env string) *App {
	return &App{Environment: env}
}

//go:embed all:dist
var dist embed.FS
var staticFS = echo.MustSubFS(dist, "dist")

// This custom Render replaces Echo's echo.Context.Render() with templ's templ.Component.Render().
func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func (a App) Start() {
	// Register the routes
	e := echo.New()
	e.HideBanner = true
	e.StaticFS("/dist", staticFS)

	e.GET("/", func(ctx echo.Context) error {
		return Render(ctx, 200, layouts.Layout("Home"))
	})
	// Start the server
	e.Logger.Fatal(e.Start(":8080"))
}
