package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mariolazzari/go-web-dev/db"
)

// list all tasks
func GetTasks(ctx *gin.Context) {
	tasks, err := db.TaskRepository.ReadTasks()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": true, "msg": "Unable to read tasks"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": false, "data": tasks})
}

// add new task and return id
func AddTask(ctx *gin.Context) {
	// request body
	var payload db.PostTaskPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": true, "msg": "Unable to read the body"})
		return
	}

	// save to db
	id, err := db.TaskRepository.AddTask(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": true, "msg": err.Error()})
		return
	}

	// response
	ctx.JSON(http.StatusOK, gin.H{"error": false, "msg": id})
}

// update existing task
func UpdateTask(ctx *gin.Context) {
	// request body
	var payload db.PatchTaskPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": true, "msg": err.Error()})
		return
	}

	// read task from db
	task, err := db.TaskRepository.ReadTask(payload.ID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": true, "msg": err.Error()})
		return
	}

	// override existing task by payload
	if payload.Title == "" {
		payload.Title = task.Title
	}

	if payload.Description == "" {
		payload.Description = task.Description
	}

	if payload.Status == "" {
		payload.Status = task.Status
	}

	// update task in db
	err = db.TaskRepository.UpdateTask(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": true, "msg": err.Error()})
		return
	}

	// response
	ctx.JSON(http.StatusOK, gin.H{"error": false, "data": payload})
}

func DeleteTask(c *gin.Context) {
	// get task id from params
	taskId := c.Param("id")
	id, err := strconv.Atoi(taskId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": true, "msg": "Invalid Id"})
		return
	}

	// check if task exists
	_, err = db.TaskRepository.ReadTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": true, "msg": "Task not found"})
		return
	}

	err = db.TaskRepository.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": true, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"error": false, "msg": fmt.Sprintf("Task with ID %d deleted successfully", id)})
}
