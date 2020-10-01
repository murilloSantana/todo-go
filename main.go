package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-go/todo"
)

func main() {
	fmt.Println("Starting the application...")
	upServer()
}

func upServer() {
	port := ":5000"

	buildRoutes()
	fmt.Printf("Starting server at port %s\n", port)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

func buildRoutes() {
	http.HandleFunc("/todo", todo.Handler)
}
