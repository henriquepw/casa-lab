package main

import (
	"github.com/henry-ns/casa-lab/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/static", "static")

	handler.IndexRouter(e)

	e.Logger.Fatal(e.Start(":3333"))
}
