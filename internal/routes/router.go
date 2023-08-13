package routes

import (
	"go-server-part3/internal/repositories"
	"go-server-part3/internal/usecases"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

func RegisterMainRoutes(e *echo.Echo, db *gorm.DB) {
	// User routers
	userRepository := repositories.NewUserRepository(db)
	userUseCase := usecases.NewUserUseCase(userRepository)
	userRoutes := NewUserRoutes(userUseCase)
	userRoutes.RegisterRoutes(e)

	// Task routers
	taskRepository := repositories.NewTaskRepository(db)
	taskUseCase := usecases.NewTaskUseCase(taskRepository)
	taskRoutes := NewTaskRoutes(taskUseCase)
	taskRoutes.RegisterRoutes(e)

}
