package routes

import (
	"go-server-part3/internal/handlers"
	"go-server-part3/internal/usecases"

	"github.com/labstack/echo"
)

type TaskRoutes struct {
	taskHandler *handlers.TaskHandler
}

func NewTaskRoutes(taskUseCase *usecases.TaskUseCase) *TaskRoutes {
	taskHandler := handlers.NewTaskHandler(taskUseCase)
	return &TaskRoutes{taskHandler}
}

func (tr *TaskRoutes) RegisterRoutes(e *echo.Echo) {
	taskGroup := e.Group("api/v1/task")
	taskGroup.POST("/create", tr.taskHandler.Create)
	taskGroup.GET("/:id", tr.taskHandler.Get)
	taskGroup.PUT("/:id", tr.taskHandler.Update)
	taskGroup.DELETE("/:id", tr.taskHandler.Delete)
}
