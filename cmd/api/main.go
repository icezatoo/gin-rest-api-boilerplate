package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/icezatoo/gin-rest-api-boilerplate/config"
	"github.com/icezatoo/gin-rest-api-boilerplate/db"
	"github.com/icezatoo/gin-rest-api-boilerplate/internal/routes"
	"github.com/icezatoo/gin-rest-api-boilerplate/pkg/httpserver"
	"github.com/sirupsen/logrus"
)

func main() {
	config := config.LoadConfig(".")

	router := SetupRouter(config)

	httpServer := httpserver.New(router, httpserver.Port(config.Post))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logrus.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		logrus.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err := httpServer.Shutdown()
	if err != nil {
		logrus.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}

func SetupRouter(config *config.Config) *gin.Engine {

	db := db.Connection(config)

	handler := gin.Default()

	routes.InitUserRoutes(db, handler)

	return handler
}
