package service

import (
	"github.com/google/uuid"
	"github.com/sirgloveface/go-iaapi/internal/model"
	"github.com/sirgloveface/go-iaapi/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) *TaskService {
	return &TaskService{repo: r}
}

func (s *TaskService) Create(title string) error {
	task := model.Task{
		ID:    uuid.New().String(),
		Title: title,
		Done:  false,
	}
	return s.repo.Create(task)
}

func (s *TaskService) List() ([]model.Task, error) {
	return s.repo.FindAll()
}

func (s *TaskService) GetByID(id string) (*model.Task, error) {
	return s.repo.FindByID(id)
}
