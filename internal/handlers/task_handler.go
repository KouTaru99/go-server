package handlers

import (
	"go-server-part3/internal/entities"
	"go-server-part3/internal/usecases"
	"net/http"

	"github.com/labstack/echo"
)

type TaskHandler struct {
	taskUseCase *usecases.TaskUseCase
}

func NewTaskHandler(taskUseCase *usecases.TaskUseCase) *TaskHandler {
	return &TaskHandler{
		taskUseCase: taskUseCase,
	}
}

func (h *TaskHandler) Create(c echo.Context) error {
	var input entities.Task
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	task, err := h.taskUseCase.CreateTask(input.TaskName, input.Title, input.Description, input.Status, input.Priority)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error creating task")
	}

	return c.JSON(http.StatusCreated, task.ID)
}

func (h *TaskHandler) Get(c echo.Context) error {
	var task = 123

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) Update(c echo.Context) error {
	var task = 123

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) Delete(c echo.Context) error {

	return c.NoContent(http.StatusNoContent)
}
