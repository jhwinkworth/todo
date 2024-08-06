package store

import (
	"errors"
	"todo/models"
)

var (
	NotFoundErr = errors.New("not found")
)

type MemStore struct {
	list map[string]models.Task
}

func NewMemStore() *MemStore {
	list := make(map[string]models.Task)
	return &MemStore{
		list,
	}
}

func (m MemStore) Add(task models.Task) error {
	m.list[task.ID] = task
	return nil
}

func (m MemStore) Get(id string) (models.Task, error) {

	if val, ok := m.list[id]; ok {
		return val, nil
	}

	return models.Task{}, NotFoundErr
}

func (m MemStore) List() (map[string]models.Task, error) {
	return m.list, nil
}

func (m MemStore) Update(id string, Task models.Task) error {

	if _, ok := m.list[id]; ok {
		m.list[id] = Task
		return nil
	}

	return NotFoundErr
}

func (m MemStore) Remove(id string) error {
	delete(m.list, id)
	return nil
}
