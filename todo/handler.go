package todo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type handler struct {
	task TaskUcase
}

func NewListHandler(t TaskUcase) Handler {
	return &handler{
		task: t,
	}
}

func (h handler) List(c *gin.Context) {
	tasks, err := h.task.list()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if tasks != nil {
		c.JSON(200, tasks)
		return
	}

	c.JSON(404, "Not found")
}

func (h handler) Create(c *gin.Context) {
	var task Task

	if err := c.ShouldBindJSON(&task); err != nil {
		fmt.Println("should")
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	err := h.task.create(task)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(201, "Created")
}
