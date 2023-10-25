package main

import (
	"log"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/printSANO/gorest-boilerplate/cmd/database"
	"github.com/printSANO/gorest-boilerplate/cmd/models"
	"github.com/printSANO/gorest-boilerplate/cmd/routes"
	"github.com/printSANO/gorest-boilerplate/config"
)

var err error

func main() {
	database.DBMain, err = database.NewSQLDB("pgx")

	if err != nil {
		log.Fatal(err)
	}

	database.DBMain.LionMigrate(&models.Example{})
	database.DBMain.LionMigrate(&models.Example2{})

	// client, ctx, err := database.ConnectNoSQLDB()

	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// 	log.Println("Disconnected from Mongo Database")
	// }()

	defer func() {
		if err = database.DBMain.Close(); err != nil {
			panic(err)
		}
		log.Println("Disconnected from SQL Database")
	}()

	r := routes.NewRouter(config.NewPortConfig())

	http.ListenAndServe(config.NewPortConfig(), r)
}
