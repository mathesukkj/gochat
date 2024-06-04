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

func SendMessage(conns chan *websocket.Conn) {

}

func HandleWs(w http.ResponseWriter, r *http.Request) {
	s := websocket.Server{Handler: websocket.Handler(NewWsServer)}
	s.ServeHTTP(w, r)
}

func InitServer(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleWs)

	go func() {
		for conn := range conns {
			for msg := range msgs {
				websocket.Message.Send(conn, msg)
				fmt.Println("loop ", msg)
			}
		}
	}()

	if err := http.ListenAndServe(port, mux); err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
