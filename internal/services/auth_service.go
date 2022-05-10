package services

import (
	dto "github.com/icezatoo/gin-rest-api-boilerplate/internal/dto/auth"
	dtoUser "github.com/icezatoo/gin-rest-api-boilerplate/internal/dto/user"
	"github.com/icezatoo/gin-rest-api-boilerplate/internal/repository"
	"github.com/jinzhu/copier"
)

type AuthService interface {
	Signin(request *dto.SigninUserRequest) (*dtoUser.UserReponse, error)
}

type authService struct {
	repository repository.AuthRepository
}

func NewAuthService(repository repository.AuthRepository) *authService {
	return &authService{repository: repository}
}

func (s *authService) Signin(request *dto.SigninUserRequest) (*dtoUser.UserReponse, error) {
	var userDto dtoUser.UserReponse
	user, err := s.repository.Signin(request)

	copier.Copy(&userDto, &user)

	return &userDto, err
}
