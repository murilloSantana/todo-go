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

func (database *DatabaseMock) create(task Task) error {
	args := database.Called(task)
	return args.Error(0)
}

func (database *DatabaseMock) list() ([]Task, error) {
	args := database.Called()
	return args.Get(0).([]Task), args.Error(1)
}

func (database *DatabaseMock) find(id string) (*Task, error) {
	args := database.Called(id)
	return args.Get(0).(*Task), args.Error(1)
}

func TestList(t *testing.T) {
	databaseMock := new(DatabaseMock)
	want := []Task{{"Task2 test", false}}

	databaseMock.On("list").Return(want, nil)

	taskUcase := NewTaskUcase(databaseMock)

	got, err := taskUcase.list()

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestListError(t *testing.T) {
	databaseMock := new(DatabaseMock)
	want := errors.New("Generic Error")

	databaseMock.On("list").Return([]Task{}, want)

	taskUcase := NewTaskUcase(databaseMock)

	resp, got := taskUcase.list()

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
		err := taskUcase.create(task)
		assert.Nil(t, err)
	}

	got, err := taskUcase.list()

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestCreateError(t *testing.T) {
	databaseMock := new(DatabaseMock)
	want := errors.New("Generic Error")

	databaseMock.On("create", mock.Anything).Return(want)

	taskUcase := NewTaskUcase(databaseMock)

	got := taskUcase.create(Task{})

	assert.Equal(t, want, got)
}

func TestFind(t *testing.T) {
	databaseMock := new(DatabaseMock)
	want := &Task{Name: "My Task", IsDone: true}

	databaseMock.On("find", mock.Anything).Return(want, nil)

	taskUcase := NewTaskUcase(databaseMock)

	got, err := taskUcase.find("122232")

	assert.Equal(t, want, got)
	assert.Nil(t, err)
}

func TestFindError(t *testing.T) {
	databaseMock := new(DatabaseMock)
	want := errors.New("Generic Error")

	databaseMock.On("find", mock.Anything).Return(&Task{}, want)

	taskUcase := NewTaskUcase(databaseMock)

	resp, got := taskUcase.find("122232")

	assert.Equal(t, want, got)
	assert.Nil(t, resp)
}