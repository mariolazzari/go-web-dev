package main

import (
	"net/http"

	"github.com/mariolazzari/go-web-dev/db"
	"github.com/mariolazzari/go-web-dev/routes"
)

func main() {
	db.InitDB()
	handler := routes.MountRoutes()

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	defer db.DB.Close()
	server.ListenAndServe()
}
