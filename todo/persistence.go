package todo

var tasks []Task

type MemoryStorage struct {}

func NewDatabase() Database {
	return MemoryStorage{}
}

func (database MemoryStorage) create(task Task) error {
	tasks = append(tasks, task)
	return nil
}

func (database MemoryStorage) list() ([]Task, error) {
	return tasks, nil
}

func (database MemoryStorage) find(id string) (*Task, error) {
	return nil, nil
}