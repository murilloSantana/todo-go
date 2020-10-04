package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"todo-go/todo"
)

func main() {
	upServer()
}

func upServer() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")

	fmt.Printf("Starting server at port %s\n", port)

	http.HandleFunc("/todo", todo.Handler)

	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil); err != nil {
		log.Fatal(err)
	}
}

func buildRoutes() {
}
