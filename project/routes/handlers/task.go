package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mariolazzari/go-web-dev/db"
)

type PostTaskPayload struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status"`
}

// save task and return id
func SaveTask(ctx *gin.Context) {
	var payload db.PostTaskPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read the body"})
		return
	}

	id, err := db.TaskRepository.SaveTaskQuery(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": true, "msg": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"error": false, "msg": id})
}
