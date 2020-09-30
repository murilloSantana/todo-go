package todo

var tasks []Task

type MemoryStorage struct {}

func NewDatabase() Database {
	return MemoryStorage{}
}

func (database MemoryStorage) persist(task Task) error {
	tasks = append(tasks, task)
	return nil
}

func (database MemoryStorage) retrieveAll() ([]Task, error) {
	return tasks, nil
}
