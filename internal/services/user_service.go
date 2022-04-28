package services

import (
	dto "github.com/icezatoo/gin-rest-api-boilerplate/internal/dto/user"
	"github.com/icezatoo/gin-rest-api-boilerplate/internal/repository"
	"github.com/jinzhu/copier"
)

type userService struct {
	repository repository.UserRepository
}

type UserService interface {
	Create(request *dto.CreateUserRequest) (*dto.UserReponse, error)
	GetUsers() ([]*dto.UserReponse, error)
	GetUserById(request *dto.RequestGetUser) (*dto.UserReponse, error)
	Update(request *dto.UpdateUserRequest) error
	Delete(request *dto.RequestDeleteUser) error
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{
		repository: repo,
	}
}

func (s *userService) GetUsers() ([]*dto.UserReponse, error) {
	var usersDto = make([]*dto.UserReponse, 0)

	users, err := s.repository.GetUsers()

	copier.Copy(&usersDto, &users)

	return usersDto, err
}

func (s *userService) Create(request *dto.CreateUserRequest) (*dto.UserReponse, error) {

	var userDto dto.UserReponse

	user, err := s.repository.Create(request)

	copier.Copy(&userDto, &user)

	return &userDto, err
}

func (s *userService) Update(request *dto.UpdateUserRequest) error {

	err := s.repository.Update(request)

	return err
}

func (s *userService) GetUserById(request *dto.RequestGetUser) (*dto.UserReponse, error) {
	var userDto dto.UserReponse

	user, err := s.repository.GetUserById(request)

	copier.Copy(&userDto, &user)

	return &userDto, err
}

func (s *userService) Delete(request *dto.RequestDeleteUser) error {
	return s.repository.Delete(request)
}
