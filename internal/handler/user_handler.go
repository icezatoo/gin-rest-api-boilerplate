package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	dto "github.com/icezatoo/gin-rest-api-boilerplate/internal/dto/user"
	customError "github.com/icezatoo/gin-rest-api-boilerplate/internal/errors"
	"github.com/icezatoo/gin-rest-api-boilerplate/internal/services"
	"github.com/icezatoo/gin-rest-api-boilerplate/pkg/formatter"
)

type UserHandler struct {
	userService services.UserService
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) GetUsers(c *gin.Context) {

	users, err := h.userService.GetUsers()

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, users)

}

func (h *UserHandler) CreateUser(c *gin.Context) {

	var request dto.CreateUserRequest

	if err := c.ShouldBind(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"errors": formatter.NewJSONFormatter().Descriptive(verr)})
			return
		}
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	user, err := h.userService.Create(&request)

	if err != nil {
		if ok := customError.IsAlredyExistsError(err); ok {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"errors": err.Error()})
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}
	c.JSON(http.StatusCreated, user)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var request dto.UpdateUserRequest
	request.ID = c.Param("id")

	if err := c.ShouldBind(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"errors": formatter.NewJSONFormatter().Descriptive(verr)})
			return
		}
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	err := h.userService.Update(&request)

	if err != nil {
		if ok := customError.IsNotFoundError(err); ok {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errors": err.Error()})
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	var request dto.RequestDeleteUser

	request.ID = c.Param("id")

	if err := c.ShouldBind(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"errors": formatter.NewJSONFormatter().Descriptive(verr)})
			return
		}
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	err := h.userService.Delete(&request)

	if err != nil {
		if ok := customError.IsNotFoundError(err); ok {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errors": err.Error()})
			return
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	var request dto.RequestGetUser
	request.ID = c.Param("id")

	if err := c.ShouldBind(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"errors": formatter.NewJSONFormatter().Descriptive(verr)})
			return
		}
		c.AbortWithStatus(http.StatusUnprocessableEntity)
		return
	}

	user, err := h.userService.GetUserById(&request)

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, user)

}
