package usecases

import (
	"go-server-part3/internal/entities"
	"go-server-part3/internal/repositories"
	"log"
)

type TaskUseCase struct {
	taskRepository repositories.TaskRepository
}

func NewTaskUseCase(taskRepository repositories.TaskRepository) *TaskUseCase {
	return &TaskUseCase{
		taskRepository: taskRepository,
	}
}

func (tc *TaskUseCase) CreateTask(task_name, title, description, status, priority string) (*entities.Task, error) {
	task := &entities.Task{
		TaskName:    task_name,
		Title:       title,
		Description: description,
	}

	// Kiểm tra và gán giá trị mặc định cho status nếu rỗng
	if status == "" {
		task.Status = "Todo"
	} else {
		task.Status = status
	}

	// Kiểm tra và gán giá trị mặc định cho priority nếu rỗng
	if priority == "" {
		task.Priority = "Low"
	} else {
		task.Priority = priority
	}
	// log.Printf("Creating task: TaskName=%s, Title=%s, Description=%s", task.TaskName, task.Title, task.Description, task.Status, task.Priority)

	createdTask, err := tc.taskRepository.Create(task)
	if err != nil {
		log.Printf("Error creating task: %v", err)
		return nil, err
	}

	log.Printf("Task created successfully: ID=%d", createdTask.ID)
	return createdTask, nil

}
