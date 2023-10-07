package database

import (
	"context"
	"database/sql"
	"log"
	"reflect"
	"time"

	"github.com/printSANO/gorest-boilerplate/config"
)

type DB struct {
	*sql.DB
}

func NewSQLDB(dbDriver string) (*DB, error) {
	db, err := connectSQLDB(dbDriver)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func connectSQLDB(dbDriver string) (*sql.DB, error) {
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

// func (D *DB) AutoMigrate(dbModel interface{}) error {

// }

func tableName(model interface{}) string {
	t := reflect.TypeOf(model)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}
