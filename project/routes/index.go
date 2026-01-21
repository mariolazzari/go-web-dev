package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mariolazzari/go-web-dev/config"
	middleware "github.com/mariolazzari/go-web-dev/middlewares"
	"github.com/mariolazzari/go-web-dev/routes/handlers"
)

func MountRoutes() *gin.Engine {
	handler := gin.Default()
	handler.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", config.Config.FEOriginURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// root handlers
	handler.GET("/", handlers.RootHandler)
	handler.NoRoute(handlers.NoRouteHandler)

	// task group
	taskRoutes := handler.Group("/tasks", middleware.AuthorizationMiddleWare())
	{
		taskRoutes.GET("/", handlers.GetTasks)
		taskRoutes.POST("/", handlers.AddTask)
		taskRoutes.PATCH("/", handlers.UpdateTask)
		taskRoutes.DELETE("/:id", handlers.DeleteTask)
	}

	// auth group
	authRoutes := handler.Group("/login")
	{
		authRoutes.GET("/google/login", handlers.HandleGoogleLogin)
		authRoutes.GET("/google/callback", handlers.HandleGoogleCallback)
	}

	return handler
}
