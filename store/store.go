package store

import "todo-app/model"

type TaskStore interface {
	Create(task model.Task) (model.Task, error)
	List() ([]model.Task, error)
	Delete(id int) error
	MarkDone(id int) error
}
