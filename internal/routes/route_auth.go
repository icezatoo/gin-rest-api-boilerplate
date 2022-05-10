package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/icezatoo/gin-rest-api-boilerplate/internal/handler"
	"github.com/icezatoo/gin-rest-api-boilerplate/internal/repository"
	"github.com/icezatoo/gin-rest-api-boilerplate/internal/services"
	"gorm.io/gorm"
)

func InitAuthRoutes(db *gorm.DB, groupRoute *gin.RouterGroup) {

	authRepository := repository.NewAuthRepository(db)
	authService := services.NewAuthService(authRepository)
	authHandler := handler.NewAuthHandler(authService)

	groupRoute.POST("/auth/login", authHandler.Signin)

}
