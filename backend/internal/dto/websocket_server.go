package dto

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/websocket"

	"github.com/mathesukkj/gochat/internal/infra/handlers"
)

type WebsocketServer struct {
	Port             string
	MessageType      string
	MessagesChan     chan Message
	ConnectedClients []*WebsocketClient
}

func NewWebsocketServer(port, messageType string) *WebsocketServer {
	return &WebsocketServer{
		Port:             port,
		MessageType:      messageType,
		MessagesChan:     make(chan Message),
		ConnectedClients: []*WebsocketClient{},
	}
}

func (s *WebsocketServer) Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandleWs(websocket.Handler(s.MessageHandler)))

	go s.BroadcastMessagesToClients()

	fmt.Println("server started in port " + s.Port + "!!")
	if err := http.ListenAndServe(s.Port, mux); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func (s *WebsocketServer) BroadcastMessagesToClients() {
	for {
		message := <-s.MessagesChan
		messageStr := message.ToString(s.MessageType)

		for _, client := range s.ConnectedClients {
			if client.Id != message.SentBy.Id && client.User != "" {
				websocket.Message.Send(client.Conn, messageStr)
			}
		}
	}
}

func (s *WebsocketServer) MessageHandler(ws *websocket.Conn) {
	client := s.AddNewClient(ws)

	for {
		message := Message{
			SentAt: time.Now(),
			SentBy: client,
		}

		if err := s.ReceiveAndSendMessageToChannel(ws, message); err != nil {
			break
		}
	}
}

func (s *WebsocketServer) AddNewClient(ws *websocket.Conn) *WebsocketClient {
	client := NewWebsocketClient(ws)
	s.ConnectedClients = append(s.ConnectedClients, client)

	return client
}

func (s *WebsocketServer) ReceiveAndSendMessageToChannel(
	conn *websocket.Conn,
	message Message,
) error {
	if err := websocket.Message.Receive(conn, &message.Message); err != nil {
		return err
	}

	s.MessagesChan <- message

	return nil
}
