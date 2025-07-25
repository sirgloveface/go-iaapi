package repository

import "github.com/sirgloveface/go-iaapi/internal/model"

type TaskRepository interface {
	Create(task model.Task) error
	FindAll() ([]model.Task, error)
	FindByID(id string) (*model.Task, error)
}
