package repository

import (
	dto "github.com/icezatoo/gin-rest-api-boilerplate/internal/dto/user"
	customError "github.com/icezatoo/gin-rest-api-boilerplate/internal/errors"
	"github.com/icezatoo/gin-rest-api-boilerplate/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(request *dto.CreateUserRequest) (*models.User, error)
	GetUsers() ([]*models.User, error)
	GetUserById(request *dto.RequestGetUser) (*models.User, error)
	Update(request *dto.UpdateUserRequest) error
	Delete(request *dto.RequestDeleteUser) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{db: db}
}

func (repo *repository) Create(request *dto.CreateUserRequest) (*models.User, error) {
	var user models.User

	resultEmail := repo.db.Select("*").Where("email = ?", request.Email).Find(&user)

	if resultEmail.RowsAffected > 0 {
		return &user, customError.AlredyExistsError("User with this email already exists")
	}

	resultUsername := repo.db.Select("*").Where("username = ?", request.Username).Find(&user)

	if resultUsername.RowsAffected > 0 {
		return &user, customError.AlredyExistsError("User with this email already exists")
	}

	user.FullName = request.FullName
	user.LastName = request.LastName
	user.Phone = request.Phone
	user.Email = request.Email
	user.Enabled = request.Enabled
	user.Password = request.Password
	user.Username = request.Username

	err := repo.db.Create(&user).Error

	return &user, err
}

func (repo *repository) GetUsers() ([]*models.User, error) {
	var users []*models.User

	err := repo.db.Find(&users).Error

	return users, err
}

func (repo *repository) GetUserById(request *dto.RequestGetUser) (*models.User, error) {
	var user models.User

	err := repo.db.First(&user, "id = ?", request.ID).Error

	return &user, err
}

func (repo *repository) Update(request *dto.UpdateUserRequest) error {
	var user models.User

	result := repo.db.Select("*").Where("id = ?", request.ID).Find(&user)

	if result.RowsAffected < 0 {
		return customError.NotFoundErrorError("User not found")
	}

	user.FullName = request.FullName
	user.LastName = request.LastName
	user.Phone = request.Phone
	user.Email = request.Email
	user.Enabled = request.Enabled

	err := repo.db.Updates(&user).Error

	return err

}

func (repo *repository) Delete(request *dto.RequestDeleteUser) error {
	result := repo.db.Delete(&models.User{}, "id = ?", request.ID)

	if result.RowsAffected < 0 {
		return customError.NotFoundErrorError("User not found")
	}

	err := repo.db.Delete(request.ID).Error

	return err

}
