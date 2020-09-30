package todo

type Task struct {
	Name   string
	IsDone bool
}

type TaskUcase interface {
	Create(task Task) error
	List() ([]Task, error)
}

type Database interface {
	persist(task Task) error
	retrieveAll() ([]Task, error)
}