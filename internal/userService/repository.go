package userService

import (
	"project/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetTasksForUser(userID uint) ([]models.Task, error)
	UpdateUserByID(id uint, user models.User) (models.User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) GetTasksForUser(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	// err := r.db.Find(&tasks).Error
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *userRepository) UpdateUserByID(id uint, user models.User) (models.User, error) {
	err := r.db.Model(&user).Where("id = ?", id).Updates(user).Error
	return user, err
}

func (r *userRepository) DeleteUserByID(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
