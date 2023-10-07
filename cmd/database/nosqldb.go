package database

import (
	"context"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/printSANO/gorest-boilerplate/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectNoSQLDB() (*mongo.Client, context.Context, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	urlDB := config.NewNOSQLDBConfig()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(urlDB))
	if err != nil {
		log.Fatal("Failed to connect to Mongo Database")
	}
	log.Println("Successfully Connected to Mongo Database")
	return client, ctx, nil
}
