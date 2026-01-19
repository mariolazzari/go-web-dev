package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	handler := gin.Default()

	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Ciao Mario"})
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	server.ListenAndServe()
}
