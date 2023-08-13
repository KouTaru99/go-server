package repositories

import (
	"go-server-part3/internal/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	GetByID(userID uint) (*entities.User, error)
	Update(user *entities.User) error
	Delete(userID uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(user *entities.User) (*entities.User, error) {
	result := r.db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (r *userRepository) GetByID(userID uint) (*entities.User, error) {
	var user entities.User
	result := r.db.First(&user, userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *userRepository) Update(user *entities.User) error {
	result := r.db.Save(user)
	return result.Error
}

func (r *userRepository) Delete(userID uint) error {
	result := r.db.Delete(&entities.User{}, userID)
	return result.Error
}
