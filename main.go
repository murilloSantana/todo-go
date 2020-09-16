package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-go/todo"
)

func main() {
	buildRoutes()
	upServer()
}

func upServer() {
	port := ":8080"

	fmt.Printf("Starting server at port %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func buildRoutes() {
	http.HandleFunc("/todo", todo.Handler)
}
