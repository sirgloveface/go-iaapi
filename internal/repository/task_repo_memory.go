package repository

import (
	"errors"

	"github.com/sirgloveface/go-iaapi/internal/model"
)

type InMemoryTaskRepository struct {
	tasks map[string]model.Task
}

func NewInMemoryTaskRepository() *InMemoryTaskRepository {
	return &InMemoryTaskRepository{
		tasks: make(map[string]model.Task),
	}
}

func (r *InMemoryTaskRepository) Create(task model.Task) error {
	r.tasks[task.ID] = task
	return nil
}

func (r *InMemoryTaskRepository) FindAll() ([]model.Task, error) {
	result := []model.Task{}
	for _, t := range r.tasks {
		result = append(result, t)
	}
	return result, nil
}

func (r *InMemoryTaskRepository) FindByID(id string) (*model.Task, error) {
	task, exists := r.tasks[id]
	if !exists {
		return nil, errors.New("task not found")
	}
	return &task, nil
}
