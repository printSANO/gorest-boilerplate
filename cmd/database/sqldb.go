package database

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"

	"github.com/printSANO/gorest-boilerplate/internal"
)

type DB struct {
	*sql.DB
}

func (d *DB) LionMigrate(dbModel interface{}) {
	t := reflect.TypeOf(dbModel)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("Model is not a struct. Migration failed")
		return
	}
	CreateSQLTable(d, t)
}

func CreateSQLTable(d *DB, t reflect.Type) {
	tableName := t.Name()
	var argsSQL []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tags := field.Tag
		columnName := tags.Get("db")
		dataType := tags.Get("dataType")
		constraint := tags.Get("constraint")
		clause := strings.TrimSpace(columnName + " " + dataType + " " + constraint)
		argsSQL = append(argsSQL, clause)
	}
	sqlArg := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", tableName, strings.Join(argsSQL, ", "))
	_, err := d.Exec(sqlArg)
	if err != nil {
		log.Println("Create Table Failed. Wrong SQL Arguments.")
		return
	}
	log.Printf("Succesfully Migrated Table Name: %s", tableName)
}

func NewSQLDB(dbDriver string) (*DB, error) {
	db, err := internal.ConnectSQLDB(dbDriver)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
