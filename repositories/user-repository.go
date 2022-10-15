package repositories

import (
	"UserManagementAPI/models"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

type UserRepository interface {
	CreateUser(models.User) error
	UpdateUser(models.User) error
	DeleteUser(int) error
	DeleteUserByEmail(string) error
	DeleteAll() error
	GetUser(int) (models.User, error)
	GetUserByEmail(string) (models.User, error)
	GetAllUsers() ([]models.User, error)
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{DB: db}
}

func (r userRepository) CreateUser(user models.User) error {
	return r.DB.Create(&user).Error
}

func (r userRepository) UpdateUser(user models.User) error {
	return r.DB.Model(&models.User{}).Where("id = ?", user.ID).Updates(&user).Error
}

func (r userRepository) DeleteUser(id int) error {
	return r.DB.Delete(&models.User{}, id).Error
}

func (r userRepository) DeleteUserByEmail(email string) error {
	return r.DB.Where("email = ", email).Delete(&models.User{}).Error
}

func (r userRepository) DeleteAll() error {
	// db.Exec("DELETE FROM users") // another simple way
	return r.DB.Where("1=1").Delete(&models.User{}).Error
}

func (r userRepository) GetUser(id int) (user models.User, err error) {
	return user, r.DB.First(&user, id).Error
}

func (r userRepository) GetUserByEmail(email string) (user models.User, err error) {
	return user, r.DB.First(&user, "email=?", email).Error
}

func (r userRepository) GetAllUsers() (users []models.User, err error) {
	return users, r.DB.Find(&users).Error
}
