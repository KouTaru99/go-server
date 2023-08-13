package repositories

import (
	"go-server-part3/internal/entities"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *entities.Task) (*entities.Task, error)
	GetByID(taskID uint) (*entities.Task, error)
	Update(task *entities.Task) error
	Delete(taskID uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

func (r *taskRepository) Create(task *entities.Task) (*entities.Task, error) {
	result := r.db.Create(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

// Delete implements TaskRepository.
func (*taskRepository) Delete(taskID uint) error {
	panic("unimplemented")
}

// GetByID implements TaskRepository.
func (*taskRepository) GetByID(taskID uint) (*entities.Task, error) {
	panic("unimplemented")
}

// Update implements TaskRepository.
func (*taskRepository) Update(task *entities.Task) error {
	panic("unimplemented")
}
