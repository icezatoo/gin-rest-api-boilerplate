package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/icezatoo/gin-rest-api-boilerplate/internal/handler"
	"github.com/icezatoo/gin-rest-api-boilerplate/internal/repository"
	"github.com/icezatoo/gin-rest-api-boilerplate/internal/services"
	"gorm.io/gorm"
)

func InitUserRoutes(db *gorm.DB, route *gin.Engine) {

	userRepository := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	groupRoute := NewRouter(route)
	groupRoute.GET("/users", userHandler.GetUsers)
	groupRoute.GET("/users/:id", userHandler.GetUserById)
	groupRoute.POST("/users", userHandler.CreateUser)
	groupRoute.PUT("/users/:id", userHandler.UpdateUser)
	groupRoute.DELETE("/users/:id", userHandler.DeleteUser)
}
