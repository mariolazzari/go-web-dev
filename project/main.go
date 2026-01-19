package main

import (
	"net/http"

	"github.com/mariolazzari/go-web-dev/routes"
)

func main() {
	handler := routes.MountRoutes()

	server := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	server.ListenAndServe()
}
