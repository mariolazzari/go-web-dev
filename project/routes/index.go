package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mariolazzari/go-web-dev/routes/handlers"
)

func MountRoutes() *gin.Engine {
	handler := gin.Default()

	handler.GET("/", handlers.RootHandler)
	handler.NoRoute(handlers.NoRouteHandler)

	handler.POST("/tasks", handlers.SaveTask)

	return handler
}
