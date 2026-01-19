package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SaveTask(ctx *gin.Context) {
	_, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
	}
	ctx.JSON(http.StatusOK, gin.H{"error": false})
}
