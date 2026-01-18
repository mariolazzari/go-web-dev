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

### Postgres

```go

```
