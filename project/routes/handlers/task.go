package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PostTaskPayload struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status"`
}

func SaveTask(ctx *gin.Context) {
	var payload PostTaskPayload

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": false, "title": payload.Title})
}
