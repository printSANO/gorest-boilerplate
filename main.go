package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/printSANO/gorest-boilerplate/cmd/database"
	"github.com/printSANO/gorest-boilerplate/cmd/models"
	"github.com/printSANO/gorest-boilerplate/cmd/routes"
	"github.com/printSANO/gorest-boilerplate/config"
	"log"
	"net/http"
)

func main() {
	db, err := database.NewSQLDB("pgx")

	if err != nil {
		log.Fatal(err)
	}

	db.LionMigrate(&models.Example{})
	db.LionMigrate(&models.Example2{})

	// client, ctx, err := database.ConnectNoSQLDB()

	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// 	log.Println("Disconnected from Mongo Database")
	// }()

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

	r := routes.NewRouter(config.NewPortConfig())

	http.ListenAndServe(config.NewPortConfig(), r)
}
