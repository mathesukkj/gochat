package main

import (
	"net/http"

	"github.com/mathesukkj/gochat/internal/infra/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HandleWs)

	err := http.ListenAndServe(":8020", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
