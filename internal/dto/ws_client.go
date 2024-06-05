package dto

import "golang.org/x/net/websocket"

type WebsocketClient struct {
	Conn *websocket.Conn
	Id   string
}
