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
	var err error

	client := dto.WebsocketClient{
		Id:   uuid.New().String(),
		Conn: ws,
	}
	clients = append(clients, &client)

	client.User, err = GetUsername(ws)
	if err != nil {
		fmt.Println(err)
	}

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

func GetUsername(ws *websocket.Conn) (username string, err error) {
	sampleUsername := "Anonymous"

	// if err := websocket.Message.Send(ws, "\033[32mWhat is your username? Please type here: \033[0m"); err != nil {
	// 	return sampleUsername, err
	// }
	//
	// if err := websocket.Message.Receive(ws, &username); err != nil {
	// 	return sampleUsername, err
	// }
	//
	// if err := websocket.Message.Send(ws, "\033[32mGreat! You can start chatting now ;)\033[0m"); err != nil {
	// 	return sampleUsername, err
	// }
	//
	// if err := websocket.Message.Send(ws, "\n"); err != nil {
	// 	return sampleUsername, err
	// }

	return sampleUsername, nil
}

func SendMessage(msgType string) {
	for {
		message := <-msgs
		messageStr := message.ToString()
		if msgType == "json" {
			messageStr = message.ToJson()
		}

		for _, client := range clients {
			if client.Id != message.SentBy.Id && client.User != "" {
				websocket.Message.Send(client.Conn, messageStr)
			}
		}
	}
}

func HandleWs(w http.ResponseWriter, r *http.Request) {
	s := websocket.Server{Handler: websocket.Handler(NewWsServer)}
	s.ServeHTTP(w, r)
}

func InitWebsocketServer(port, msgType string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleWs)

	go SendMessage(msgType)

	if err := http.ListenAndServe(port, mux); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
