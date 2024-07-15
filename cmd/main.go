package main

import (
	"github.com/henry-ns/casa-lab/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/websocket"
)

func main() {
	e := echo.New()

	e.Static("/static", "static")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/reload", func(c echo.Context) error {
		websocket.Handler(func(ws *websocket.Conn) {
			defer ws.Close()

			err := websocket.Message.Send(ws, "reload")
			if err != nil {
				c.Logger().Error(err)
			}

			msg := ""
			err = websocket.Message.Receive(ws, &msg)
			if err != nil {
				c.Logger().Error(err)
			}
		}).ServeHTTP(c.Response(), c.Request())

		return nil
	})

	handler.IndexRouter(e)
	e.Logger.Fatal(e.Start("localhost:3333"))
}
