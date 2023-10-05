package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/joho/godotenv"
)

func ConnectSQLDB(dbType string, migration ...bool) (*sql.DB, error) {
	migrate := false
	if len(migration) > 0 {
		migrate = migration[0]
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	urlDB := "postgres://postgres:postgres@localhost:5432/boilerplate"
	// urlDB := os.Getenv("SQL_DB_URL")
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

	log.Println("Database Connection Success!")
	if migrate {
		log.Println("Migration Success!")
	}
	return db, nil
}
