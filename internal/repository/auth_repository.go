package repository

import (
	dto "github.com/icezatoo/gin-rest-api-boilerplate/internal/dto/auth"
	customError "github.com/icezatoo/gin-rest-api-boilerplate/internal/errors"
	"github.com/icezatoo/gin-rest-api-boilerplate/internal/models"
	"github.com/icezatoo/gin-rest-api-boilerplate/pkg/bcrypt"
	"gorm.io/gorm"
)

type AuthRepository interface {
	Signin(request *dto.SigninUserRequest) (*models.User, error)
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &repository{db: db}
}

func (repo *repository) Signin(request *dto.SigninUserRequest) (*models.User, error) {

	var user models.User

	err := repo.db.First(&user, "username = ?", request.Username).Error

	err = bcrypt.ComparePassword(request.Password, user.Password)

	if err != nil {
		return &user, customError.AuthFailedError("Incorrect Username or Password")
	}

	return &user, nil

}
