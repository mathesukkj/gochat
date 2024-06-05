package handlers

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/websocket"

	"github.com/mathesukkj/gochat/internal/dto"
)

var conns = make(chan *websocket.Conn)
var msgs = make(chan string)

func NewWsServer(ws *websocket.Conn) {
	conns <- ws
	for {
		message := dto.Message{
			SentAt: time.Now(),
		}

		if err := websocket.Message.Receive(ws, &message.Message); err != nil {
			fmt.Println(err)
			break
		}

		msgs <- message.ToString()
	}
}

func SendMessage() {
	for {
		fmt.Println(<-conns)
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

	go func() {
		for {
			fmt.Println(<-msgs)
		}
	}()

	if err := http.ListenAndServe(port, mux); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
