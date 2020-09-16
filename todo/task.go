package todo

import "fmt"

type Task struct {
	Name   string
	IsDone bool
}

func (task Task) Create() {
	err := database.persist(task)

	if err != nil {
		fmt.Println(err)
	}
}

func List() []Task {
	tasks, err := database.retrieveAll()

	if err != nil {
		fmt.Println(err)
	}

	return tasks
}
