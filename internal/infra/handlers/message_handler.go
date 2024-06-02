package handlers

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

func NewWsServer(ws *websocket.Conn) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(ws.Request().Body)
	fmt.Println(buf)
	io.Copy(buf, ws)
}

func HandleWs(w http.ResponseWriter, req *http.Request) {
	s := websocket.Server{Handler: websocket.Handler(NewWsServer)}
	s.ServeHTTP(w, req)
}
