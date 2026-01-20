package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RootHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Ciao Mario"})
}

func NoRouteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"message": "Route not found"})
}
