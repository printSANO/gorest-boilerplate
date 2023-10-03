package main

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/printSANO/gorest-boilerplate/database"
)

func main() {
	// dbURL := "postgres://postgres:postgres@localhost:5432/boilerplate"
	db, err := database.ConnectDB("pgx")

	if err != nil {
		log.Fatal(err)
	}

	var greeting string
	err = db.QueryRow("select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}
	fmt.Println(greeting)

	defer db.Close()
}
