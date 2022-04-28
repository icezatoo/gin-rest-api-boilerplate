package routes

import (
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/heptiolabs/healthcheck"
)

func NewRouter(handler *gin.Engine) *gin.RouterGroup {
	// Options
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

	// Swagger
	// swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	// handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	health := healthcheck.NewHandler()

	health.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))

	handler.GET("/healthz", gin.WrapF(health.LiveEndpoint))

	// Prometheus metrics
	// handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	groupRoute := handler.Group("/api/v1")

	return groupRoute

}
