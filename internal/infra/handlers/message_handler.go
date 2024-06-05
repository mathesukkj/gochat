package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/net/websocket"

	"github.com/mathesukkj/gochat/internal/dto"
)

var clients = make([]*dto.WebsocketClient, 0)
var msgs = make(chan dto.Message)

func NewWsServer(ws *websocket.Conn) {
	client := dto.WebsocketClient{
		Id:   uuid.New().String(),
		Conn: ws,
	}
	clients = append(clients, &client)

	for {
		message := dto.Message{
			SentAt: time.Now(),
			SentBy: client,
		}

		if err := websocket.Message.Receive(ws, &message.Message); err != nil {
			fmt.Println(err)
			break
		}

		msgs <- message
	}
}

func SendMessage() {
	for {
		message := <-msgs
		for _, client := range clients {
			if client.Id != message.SentBy.Id {
				websocket.Message.Send(client.Conn, message.ToString())
			}
		}
	}
}

func HandleWs(w http.ResponseWriter, r *http.Request) {
	s := websocket.Server{Handler: websocket.Handler(NewWsServer)}
	s.ServeHTTP(w, r)
}

func InitServer(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleWs)

	go SendMessage()

	if err := http.ListenAndServe(port, mux); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
