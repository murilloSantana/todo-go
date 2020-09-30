package todo

type Task struct {
	Name   string
	IsDone bool
}

type TaskUcase interface {
	create(task Task) error
	list() ([]Task, error)
	find(id string) (*Task, error)
}

type Database interface {
	create(task Task) error
	list() ([]Task, error)
	find(id string) (*Task, error)
}