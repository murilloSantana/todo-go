package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

	if err := http.ListenAndServe(os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}

func buildRoutes() {
	http.HandleFunc("/todo", todo.Handler)
}
