package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/icezatoo/gin-rest-api-boilerplate/internal/handler"
	"github.com/icezatoo/gin-rest-api-boilerplate/internal/repository"
	"github.com/icezatoo/gin-rest-api-boilerplate/internal/services"
	auth "github.com/icezatoo/gin-rest-api-boilerplate/pkg/jwt"
	"gorm.io/gorm"
)

func InitUserRoutes(db *gorm.DB, groupRoute *gin.RouterGroup) {

	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	groupRoute.GET("/users", auth.AuthorizeJWT("JWT_SECRET"), userHandler.GetUsers)
	groupRoute.GET("/users/:id", auth.AuthorizeJWT("JWT_SECRET"), userHandler.GetUserById)
	groupRoute.POST("/users", userHandler.CreateUser)
	groupRoute.PUT("/users/:id", auth.AuthorizeJWT("JWT_SECRET"), userHandler.UpdateUser)
	groupRoute.DELETE("/users/:id", auth.AuthorizeJWT("JWT_SECRET"), userHandler.DeleteUser)

}
