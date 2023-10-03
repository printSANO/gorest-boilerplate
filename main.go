package main

import (
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/printSANO/gorest-boilerplate/database"
)

func main() {
	db, err := database.ConnectSQLDB("pgx")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	var greeting string
	err = db.QueryRow("select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}
	fmt.Println(greeting)
}
