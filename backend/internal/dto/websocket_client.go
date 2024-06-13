package dto

import (
	"github.com/google/uuid"
	"golang.org/x/net/websocket"
)

type WebsocketClient struct {
	Conn *websocket.Conn
	User string
	Id   string
}

func NewWebsocketClient(ws *websocket.Conn) *WebsocketClient {
	return &WebsocketClient{
		Id:   uuid.New().String(),
		Conn: ws,
		User: GetClientUsername(ws),
	}
}

// TODO: redo it in json
func GetClientUsername(ws *websocket.Conn) string {
	sampleUsername := "Anonymous"
	return sampleUsername
}
