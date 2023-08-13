package usecases

import "go-server-part3/internal/repositories"

type TaskUseCase struct {
	taskRepository repositories.TaskRepository
}

func NewTaskUseCase(taskRepository repositories.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		taskRepository: taskRepository,
	}
}
