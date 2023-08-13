package handlers

import (
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
	var task = 123

	return c.JSON(http.StatusCreated, task)
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
