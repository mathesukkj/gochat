package main

import (
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

func NewWsServer(ws *websocket.Conn) {
	io.Copy(ws, ws)
}

func HandleWs(w http.ResponseWriter, req *http.Request) {
	s := websocket.Server{Handler: websocket.Handler(NewWsServer)}
	s.ServeHTTP(w, req)
}

func main() {
	http.HandleFunc("/", HandleWs)

	err := http.ListenAndServe(":8020", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
