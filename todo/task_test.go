package todo

import "testing"

func TestSave(t *testing.T) {
	task := Task{
		Name:   "Task test",
		IsDone: true,
	}

	task.Save()
}
