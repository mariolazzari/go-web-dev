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
