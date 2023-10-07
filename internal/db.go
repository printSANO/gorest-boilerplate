package internal

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/printSANO/gorest-boilerplate/config"
)

func ConnectSQLDB(dbDriver string) (*sql.DB, error) {
	urlDB := config.NewSQLDBConfig()
	db, err := sql.Open(dbDriver, urlDB)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	log.Println("Database Connection Success!")
	return db, nil
}
