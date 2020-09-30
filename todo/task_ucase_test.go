package todo

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/stretchr/testify/mock"
)

type DatabaseMock struct {
	mock.Mock
}

func (database *DatabaseMock) persist(task Task) error {
	args := database.Called(task)
	return args.Error(0)
}

func (database *DatabaseMock) retrieveAll() ([]Task, error) {
	args := database.Called()
	return args.Get(0).([]Task), args.Error(1)
}

func TestList(t *testing.T) {
	databaseMock := new(DatabaseMock)

	want := []Task{{"Task2 test", false}}

	databaseMock.On("retrieveAll").Return(want, nil)

	taskUcase := NewTaskUcase(databaseMock)

	got, err := taskUcase.List()

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestListError(t *testing.T) {
	databaseMock := new(DatabaseMock)
	want := errors.New("Generic Error")

	databaseMock.On("retrieveAll").Return([]Task{}, want)

	taskUcase := NewTaskUcase(databaseMock)

	resp, got := taskUcase.List()

	assert.Nil(t, resp)
	assert.Equal(t, want, got)
}

func TestCreate(t *testing.T) {
	want := []Task{
		{"Task test", true},
		{"Task2 test", false},
	}

	taskUcase := NewTaskUcase(NewDatabase())

	for _, task := range want {
		taskUcase.Create(task)
	}

	got, err := taskUcase.List()

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestCreateError(t *testing.T) {
	databaseMock := new(DatabaseMock)

	want := errors.New("Generic Error")

	databaseMock.On("persist", mock.Anything).Return(want)

	taskUcase := NewTaskUcase(databaseMock)

	got := taskUcase.Create(Task{})

	assert.Equal(t, want, got)
}
