package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func ConnectDb() {
	ctx := context.Background()

	var err error
	err = godotenv.Load()
	if err != nil {
		panic(err)
	}
	connStr := os.Getenv("DB_URL")
	Pool, err = pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}
	fmt.Println("Database Connected successfully")
}

func CloseConnection() {
	Pool.Close()
}
