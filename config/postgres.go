package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/joho/godotenv/autoload"
)

const DBTimeout = time.Second * 5

func PsqlSetup() *pgxpool.Pool {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), DBTimeout)
	defer cancel()

	dsn := os.Getenv("DSN")

	pgConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Panic(err)
	} 


	db, err := pgxpool.ConnectConfig(ctx, pgConfig)
	if err != nil {
		log.Panic(err)
	}

	return db
}