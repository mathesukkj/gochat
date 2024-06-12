package main

import (
	"github.com/mathesukkj/gochat/internal/infra/handlers"
)

func main() {
	handlers.InitWebsocketServer(":8020", "json")
}
