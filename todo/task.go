package todo

import (
	"fmt"
)

type Task struct {
	Name   string
	IsDone bool
}

func (task Task) Save() {
	fmt.Println(task)
}
