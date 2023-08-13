package repositories

import (
	"go-server-part3/internal/entities"
	"go-server-part3/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	GetByID(userID uint) (*entities.User, error)
	Update(user *entities.User) error
	Delete(userID uint) error
	Search(request models.SearchRequest) ([]*entities.User, error)
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

func (r *userRepository) Search(request models.SearchRequest) ([]*entities.User, error) {
	// Xây dựng truy vấn tìm kiếm và lọc dữ liệu
	query := r.db
	if request.SearchKey != "" {
		query = query.Where("username LIKE ?", "%"+request.SearchKey+"%").
			Or("email LIKE ?", "%"+request.SearchKey+"%")
	}
	for field, value := range request.Filters {
		query = query.Where(field+" = ?", value)
	}
	if request.SortBy != "" {
		query = query.Order(request.SortBy + " " + request.SortOrder)
	}

	// Phân trang dữ liệu
	offset := (request.Page - 1) * request.PageSize
	var users []*entities.User
	result := query.Offset(offset).Limit(request.PageSize).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
