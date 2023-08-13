package usecases

import (
	"go-server-part3/internal/entities"
	"go-server-part3/internal/models"
	"go-server-part3/internal/repositories"
)

type UserUseCase struct {
	userRepository repositories.UserRepository
}

func NewUserUseCase(userRepository repositories.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepository: userRepository,
	}
}

func (uc *UserUseCase) CreateUser(username, email string) (*entities.User, error) {
	user := &entities.User{
		Username: username,
		Email:    email,
	}
	return uc.userRepository.Create(user)
}

func (uc *UserUseCase) GetUserByID(userID uint) (*entities.User, error) {
	return uc.userRepository.GetByID(userID)
}

func (uc *UserUseCase) UpdateUser(user *entities.User) error {
	return uc.userRepository.Update(user)
}

func (uc *UserUseCase) DeleteUser(userID uint) error {
	return uc.userRepository.Delete(userID)
}

func (uc *UserUseCase) SearchUsers(request models.SearchRequest) ([]*entities.User, error) {
	// Gọi hàm từ repository để thực hiện tìm kiếm và lọc người dùng
	return uc.userRepository.Search(request)
}
