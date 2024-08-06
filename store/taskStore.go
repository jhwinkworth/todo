package store

import "todo/models"

type TaskStore interface {
	Add(task models.Task) error
	Get(id string) (models.Task, error)
	Update(id string, task models.Task) error
	List() (map[string]models.Task, error)
	Remove(id string) error
}
