package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mariolazzari/go-web-dev/postgres/config"
	"github.com/mariolazzari/go-web-dev/postgres/db"
)

func main() {
	handler := gin.Default()

	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Ciao Mario",
		})
	})

	db.InitDB()
	defer db.DB.Close()

	server := &http.Server{
		Addr:    config.Config.AppPort,
		Handler: handler,
	}

	server.ListenAndServe()
}
