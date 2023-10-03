package database

import (
	"context"
	"database/sql"
	"log"
	"time"
)

func ConnectSQLDB(dbType string) (*sql.DB, error) {
	urlDB := "postgres://postgres:postgres@localhost:5432/boilerplate"
	// urlDB := os.Getenv("DB_URL")
	db, err := sql.Open(dbType, urlDB)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return db, nil
}
