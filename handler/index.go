package handler

import (
	"net/http"

	"github.com/henry-ns/casa-lab/view/layout"
	"github.com/labstack/echo/v4"
)

func IndexRouter(app *echo.Echo) {
	app.GET("/", func(ctx echo.Context) error {
		return Render(ctx, http.StatusOK, layout.Hello())
	})
}
