package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	dto "github.com/icezatoo/gin-rest-api-boilerplate/internal/dto/auth"
	customError "github.com/icezatoo/gin-rest-api-boilerplate/internal/errors"
	"github.com/icezatoo/gin-rest-api-boilerplate/internal/services"
	"github.com/icezatoo/gin-rest-api-boilerplate/pkg/formatter"
	auth "github.com/icezatoo/gin-rest-api-boilerplate/pkg/jwt"
)

type AuthHandler struct {
	authService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Signin(c *gin.Context) {
	var request dto.SigninUserRequest

	if err := c.ShouldBind(&request); err != nil {
		var verr validator.ValidationErrors
		if errors.As(err, &verr) {
			c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"errors": formatter.NewJSONFormatter().Descriptive(verr)})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": "Something went wrong"})
		return
	}

	user, err := h.authService.Signin(&request)

	if err != nil {
		if ok := customError.IsAuthFailedError(err); ok {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
			return
		}
	}

	token, err := auth.Sign(map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"fullName": user.FullName,
		"email":    user.Email,
		"lastName": user.LastName,
	}, "JWT_SECRET", time.Duration(10))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}
