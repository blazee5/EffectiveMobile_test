package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

func New(ctx context.Context) *pgxpool.Pool {
	db, err := pgxpool.New(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE")))

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	if err := db.Ping(ctx); err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	return db
}
