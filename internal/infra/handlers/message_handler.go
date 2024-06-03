package handlers

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/websocket"

	"github.com/mathesukkj/gochat/internal/dto"
)

func NewWsServer(ws *websocket.Conn) {
	for {
		message := dto.Message{
			SentAt: time.Now(),
		}

		if err := websocket.Message.Receive(ws, &message.Message); err != nil {
			fmt.Println(err)
			break
		}

		if err := websocket.Message.Send(ws, message.ToString()); err != nil {
			fmt.Println(err)
			break
		}
	}
}

func HandleWs(w http.ResponseWriter, req *http.Request) {
	s := websocket.Server{Handler: websocket.Handler(NewWsServer)}
	s.ServeHTTP(w, req)
}
