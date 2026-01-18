package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mariolazzari/go-web-dev/env/config"
)

func main() {
	handler := gin.Default()

	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Ciao Mario",
		})
	})

	server := &http.Server{
		Addr:    config.Config.AppPort,
		Handler: handler,
	}

	server.ListenAndServe()
}
