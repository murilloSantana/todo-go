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
	fmt.Println("Starting the application...")
	upServer()
}

func upServer() {
	godotenv.Load()
	port := os.Getenv("PORT")

	buildRoutes()
	fmt.Printf("Starting server at port %s\n", port)

	if err := http.ListenAndServe(fmt.Sprintf(port), nil); err != nil {
		log.Fatal(err)
	}
}

func buildRoutes() {
	http.HandleFunc("/todo", todo.Handler)
}
