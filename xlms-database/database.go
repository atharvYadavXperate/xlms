package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func ConnectDb() {
	ctx := context.Background()

	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system env")
	}

	connStr := os.Getenv("DB_URL")
	if connStr == "" {
		log.Fatal("DB_URL is not set")
	}

	var err error
	Pool, err = pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
	}

	// 🔥 IMPORTANT: verify connection
	if err = Pool.Ping(ctx); err != nil {
		log.Fatalf("Database ping failed: %v\n", err)
	}

	log.Println("Database Connected successfully")
}

func CloseConnection() {
	if Pool != nil {
		Pool.Close()
	}
}
