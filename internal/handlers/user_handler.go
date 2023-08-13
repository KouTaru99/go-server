package handlers

import (
	"go-server-part3/internal/usecases"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type UserHandler struct {
	userUseCase *usecases.UserUseCase
}

func NewUserHandler(userUseCase *usecases.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (h *UserHandler) CreateUser(c echo.Context) error {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	user, err := h.userUseCase.CreateUser(input.Username, input.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error creating user")
	}

	return c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) GetUser(c echo.Context) error {
	userIDStr := c.Param("id")

	// Chuyển đổi userID sang dạng số
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}
	// Xử lý lỗi nếu không thể chuyển đổi
	// ...
	user, err := h.userUseCase.GetUserByID(uint(userID))
	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {

	userIDStr := c.Param("id")
	// Chuyển đổi userID sang dạng số
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}
	// Xử lý lỗi nếu không thể chuyển đổi
	// ...

	user, err := h.userUseCase.GetUserByID(uint(userID))
	if err != nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	if err := c.Bind(&input); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid input")
	}

	user.Username = input.Username
	user.Email = input.Email

	if err := h.userUseCase.UpdateUser(user); err != nil {
		return c.JSON(http.StatusInternalServerError, "Error updating user")
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	userIDStr := c.Param("id")
	// Chuyển đổi userID sang dạng số
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}
	// Xử lý lỗi nếu không thể chuyển đổi
	// ...

	if err := h.userUseCase.DeleteUser(uint(userID)); err != nil {
		return c.JSON(http.StatusInternalServerError, "Error deleting user")
	}

	return c.NoContent(http.StatusNoContent)
}
