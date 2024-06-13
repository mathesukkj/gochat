package handlers

import (
	"net/http"

	"golang.org/x/net/websocket"
)

func HandleWs(messageHandler websocket.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := websocket.Server{Handler: websocket.Handler(messageHandler)}
		s.ServeHTTP(w, r)
	}
}
