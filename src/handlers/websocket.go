package handlers

import (
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type responder interface{
	GetResponse() string
}

type WsHandlerFnT func(ws *websocket.Conn, respchan chan responder, errchan chan error, sleep int) error

func HandleWSConnection(c echo.Context, WsHandlerFn WsHandlerFnT, respchan chan responder, errchan chan error, sleep int) error {
	var err error
	websocket.Handler(func(ws *websocket.Conn){
		defer ws.Close()
		for err == nil {
			err = WsHandlerFn(ws, respchan, errchan, sleep)
		}
	}).ServeHTTP(c.Response(), c.Request())
	if err.Error() != "done" {
		return err
	}
	return nil
}

