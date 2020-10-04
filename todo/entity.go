package todo

import "github.com/gin-gonic/gin"

type Task struct {
	Name   string `json: "task_name"`
	IsDone bool   `json:"is_done"`
}

type TaskUcase interface {
	create(task Task) error
	list() ([]Task, error)
	find(id string) (*Task, error)
}

type Handler interface {
	List(c *gin.Context)
	Create(c *gin.Context)
}

type Database interface {
	create(task Task) error
	list() ([]Task, error)
	find(id string) (*Task, error)
}
