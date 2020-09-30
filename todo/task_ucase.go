package todo

type ucase struct {
	db Database
}

func NewTaskUcase(db Database) TaskUcase {
	return ucase{
		db: db,
	}
}

func (u ucase) Create(task Task) error {
	err := u.db.persist(task)

	if err != nil {
		return err
	}

	return nil
}

func (u ucase) List() ([]Task, error) {
	tasks, err := u.db.retrieveAll()

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
