package main

import (
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/printSANO/gorest-boilerplate/database"
)

func main() {
	db, err := database.ConnectSQLDB("pgx")

	if err != nil {
		log.Fatal(err)
	}

	client, ctx, err := database.ConnectNoSQLDB()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
		log.Println("Disconnected from Mongo Database")
	}()

	defer func() {
		if err = db.Close(); err != nil {
			panic(err)
		}
		log.Println("Disconnected from SQL Database")
	}()

	var greeting string
	err = db.QueryRow("select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Fatalf("QueryRow failed: %v\n", err)
	}
	log.Println(greeting)
}
