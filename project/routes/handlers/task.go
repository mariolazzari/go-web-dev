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

// list all tasks
func GetTasks(ctx *gin.Context) {
	tasks, err := db.TaskRepository.ReadTasks()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": true, "msg": "Unable to read tasks"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": false, "data": tasks})
}

// save task and return id
func SaveTask(ctx *gin.Context) {
	// request body
	var payload db.PostTaskPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": true, "msg": "Unable to read the body"})
		return
	}

	// save to db
	id, err := db.TaskRepository.SaveTask(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": true, "msg": err.Error()})
		return
	}

	// response
	ctx.JSON(http.StatusOK, gin.H{"error": false, "msg": id})
}
