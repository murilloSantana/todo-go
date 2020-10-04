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

	d := todo.NewDatabase()
	t := todo.NewTaskUcase(d)
	td := todo.NewListHandler(t)

	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/todo", td.List)
		v1.POST("/todo", td.Create)
	}


	if err := r.Run(fmt.Sprintf(":%v", os.Getenv("PORT"))); err != nil {
		log.Fatal(err)
	}
}
