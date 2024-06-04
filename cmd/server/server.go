package main

import (
	"github.com/mathesukkj/gochat/internal/infra/handlers"
)

func main() {
	handlers.InitServer(":8020")
}
