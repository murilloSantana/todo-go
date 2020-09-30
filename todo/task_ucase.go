package todo

type ucase struct {
	db Database
}

func NewTaskUcase(db Database) TaskUcase {
	return ucase{
		db: db,
	}
}

func (u ucase) create(task Task) error {
	err := u.db.create(task)

	if err != nil {
		return err
	}

	return nil
}

func (u ucase) list() ([]Task, error) {
	tasks, err := u.db.list()

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (u ucase) find(id string) (*Task, error){
	task, err := u.db.find(id)

	if err != nil {
		return nil, err
	}

	return task, nil
}