package todo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type DatabaseMock struct {
	mock.Mock
}

func (database *DatabaseMock) retrieveAll() ([]Task, error) {
	args := database.Called()
	return args.Get(0).([]Task), args.Error(0)
}

func TestCreate(t *testing.T) {
	want := []Task{
		{"Task test", true},
		{"Task2 test", false},
	}

	taskTable := []Task{
		{"Task test", true},
		{"Task2 test", false},
	}

	for _, task := range taskTable {
		task.Create()
	}

	got := List()

	assert.Equal(t, want, got)
}
func TestList(t *testing.T) {
	databaseMock := new(DatabaseMock)

	databaseMock.On("retrieveAll").Return(nil, nil)

	fmt.Println(List())
}
