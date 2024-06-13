package main

import "github.com/mathesukkj/gochat/internal/dto"

func main() {
	server := dto.NewWebsocketServer(":8000", "json")
	server.Serve()
}
