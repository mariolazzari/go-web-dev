# Golang Web Development: Create Powerful Servers with Golang

## Intro

[Go](https://go.dev/)

## First server

```sh
brew install go
mkdir 02-first
cd 02-first
go mod init github.com/mariolazzari/go-web-dev/02-first
```

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ciao Mario")
	})

	server.ListenAndServe()
}
```

## Gin Gonic

[Gin Gonic](https://gin-gonic.com/)
[Github](https://github.com/gin-gonic/gin)

```sh
mkdir 03-gin && cd 03-gin
go mod init github.com/mariolazzari/go-web-dev/gin
go get -u github.com/gin-gonic/gin
```

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	handler := gin.Default()

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Ciao Mario",
		})
	})

	server.ListenAndServe()
}
```

## Enviroment setup

[godotenv](https://github.com/joho/godotenv)

```sh
mkdir 04-env && cd 04-env
go mod init github.com/mariolazzari/go-web-dev/env
go get -u github.com/gin-gonic/gin
go get github.com/joho/godotenv
```

```go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type envConfig struct {
	AppPort string
}

func (e *envConfig) LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Env file not loaded")
	}

	e.AppPort = loadString("APP_PORT", "8080")
}

var Config envConfig

func loadString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("Missing var: %s", key)
		return fallback
	}

	return val
}

func init() {
	Config.LoadConfig()
}
```

## Postgres

[pgx](https://github.com/jackc/pgx)

```sh
go mod init github.com/mariolazzari/go-web-dev/postgres
go get github.com/jackc/pgx/v5
```

```go
package db

import (
	"context"
	"database/sql"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
	// read env
	err := godotenv.Load()
	if err != nil {
		log.Panic("Env file not loaded")
		os.Exit(1)
	}

	// db connection
	DB, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to connect: %v\n", err)
		os.Exit(1)
	}

	// check connection
	err = DB.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
		os.Exit(1)
	}
	log.Println("Connected to DB")
}
```

## Air

```sh
go install github.com/air-verse/air@latest
air init
air
```

## POST request

### Data validation

```go
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
```

### Data migration

[migrate](https://github.com/golang-migrate/migrate)

```sh
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
migrate create -ext sql -dir db/migrations -seq create_task_table
```

```sql
DROP TABLE IF EXISTS tasks;

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) DEFAULT 'pending',
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
```

### Saving data

```go
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
```

## Code hygiene

### Cleaning code

```go
package db

import "context"

type Task struct{}

var TaskRepository = Task{}

type PostTaskPayload struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status"`
}

func (t *Task) SaveTaskQuery(payload PostTaskPayload) (int, error) {
	var id int

	query := `Insert into tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id;`

	err := DB.QueryRow(context.Background(), query, payload.Title, payload.Description, payload.Status).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
```

### Makefile

```sh
include .env

MIGRATIONS_PATH = ./db/migrations

.PHONY: migrate-create
migrate-create:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_PATH) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_PATH) down $(filter-out $@,$(MAKECMDGOALS))
```

### Route not found

```go
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mariolazzari/go-web-dev/routes/handlers"
)

func MountRoutes() *gin.Engine {
	handler := gin.Default()

	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Ciao Mario"})
	})

	handler.POST("/tasks", handlers.SaveTask)
	handler.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Route not found"})
	})

	return handler
}
```

## Read, update and delete APIs

```go
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

```
