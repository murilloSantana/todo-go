package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
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

	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
	r.GET("/todo", todo.Handler)

	if err := r.Run(fmt.Sprintf(":%v", os.Getenv("PORT"))); err != nil {
		log.Fatal(err)
	}
}