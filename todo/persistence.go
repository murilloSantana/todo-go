package todo

var database MemoryStorage

type Database interface {
	persist(task Task) error
	retrieveAll() ([]Task, error)
}

type MemoryStorage struct {
	Tasks []Task
}

func init() {
	database = MemoryStorage{}
}

func (database *MemoryStorage) persist(task Task) error {
	database.Tasks = append(database.Tasks, task)
	return nil
}

func (database *MemoryStorage) retrieveAll() ([]Task, error) {
	return database.Tasks, nil
}
