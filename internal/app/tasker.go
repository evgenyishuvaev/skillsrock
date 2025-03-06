package app

import (
	"skillsrock/internal/storage"
)

type Tasker struct {
	storage Storage
}

type Storage interface {
	CreateTask(task storage.Task) error
	UpdateTask(id string, event storage.Task) error
	DeleteTask(id string) error
	ListTask() ([]storage.Task, error)
}

func (t *Tasker) CreateTask(task storage.Task) error {
	err := t.storage.CreateTask(task)
	if err != nil {
		return err
	}
	return nil
}

func (t *Tasker) UpdateTask(id string, task storage.Task) error {
	err := t.storage.UpdateTask(id, task)
	if err != nil {
		return err
	}
	return nil
}

func (t *Tasker) DeleteTask(id string) error {
	err := t.storage.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}

func (t *Tasker) ListTask() ([]storage.Task, error) {
	res, err := t.storage.ListTask()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewTasker(storage Storage) (*Tasker, error) {
	return &Tasker{storage: storage}, nil
}
