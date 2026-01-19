package handlers

import (
	"context"
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
	var payload PostTaskPayload

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	var id int
	query := `Insert into tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id;`

	err = db.DB.QueryRow(context.Background(), query, payload.Title, payload.Description, payload.Status).Scan(&id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": true, "message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": false, "id": id})
}
