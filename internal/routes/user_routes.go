package routes

import (
	"go-server-part3/internal/handlers"
	"go-server-part3/internal/usecases"

	"github.com/labstack/echo"
)

type UserRoutes struct {
	userHandler *handlers.UserHandler
}

func NewUserRoutes(userUseCase *usecases.UserUseCase) *UserRoutes {
	userHandler := handlers.NewUserHandler(userUseCase)
	return &UserRoutes{userHandler}
}

func (ur *UserRoutes) RegisterRoutes(e *echo.Echo) {
	userGroup := e.Group("api/v1/user")
	userGroup.POST("/create", ur.userHandler.CreateUser)
	userGroup.GET("/:id", ur.userHandler.GetUser)
	userGroup.PUT("/:id", ur.userHandler.UpdateUser)
	userGroup.DELETE("/:id", ur.userHandler.DeleteUser)
	userGroup.POST("/search", ur.userHandler.SearchUser)
}
