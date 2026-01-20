package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mariolazzari/go-web-dev/routes/handlers"
)

func MountRoutes() *gin.Engine {
	handler := gin.Default()

	// root handlers
	handler.GET("/", handlers.RootHandler)
	handler.NoRoute(handlers.NoRouteHandler)
	// task group
	taskRoutes := handler.Group("/tasks")
	{
		taskRoutes.GET("/", handlers.GetTasks)
		taskRoutes.POST("/", handlers.AddTask)
		taskRoutes.PATCH("/", handlers.UpdateTask)
		taskRoutes.DELETE("/:id", handlers.DeleteTask)
	}

	return handler
}
