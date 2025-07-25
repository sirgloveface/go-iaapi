package repository

import (
	"github.com/sirgloveface/go-iaapi/internal/model"
	"gorm.io/gorm"
)

type PostgresTaskRepository struct {
	db *gorm.DB
}

func NewPostgresTaskRepository(db *gorm.DB) *PostgresTaskRepository {
	return &PostgresTaskRepository{db: db}
}

func (r *PostgresTaskRepository) Create(task model.Task) error {
	return r.db.Create(&task).Error
}

func (r *PostgresTaskRepository) FindAll() ([]model.Task, error) {
	var tasks []model.Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *PostgresTaskRepository) FindByID(id string) (*model.Task, error) {
	var task model.Task
	err := r.db.First(&task, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}
