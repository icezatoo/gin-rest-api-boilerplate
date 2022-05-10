package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/heptiolabs/healthcheck"
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

	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())
	handler.Use(helmet.Default())
	handler.Use(gzip.Gzip(gzip.DefaultCompression))

	handler.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	if config.Environment != "production" && config.Environment != "test" {
		gin.SetMode(gin.DebugMode)
	} else if config.Environment == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	health := healthcheck.NewHandler()

	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))

	handler.GET("/healthz", gin.WrapF(health.LiveEndpoint))

	groupRoute := handler.Group("/api/v1")
	routes.InitUserRoutes(db, groupRoute)
	routes.InitAuthRoutes(db, groupRoute)

	return handler
}
