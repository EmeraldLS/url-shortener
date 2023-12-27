package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var URL_COLLECTIONS *mongo.Collection

func dbClient() (*mongo.Client, error) {
	var conn_string = os.Getenv("URL_SHORTENER_DB_CONNECTION_STRING")
	clOptions := options.Client().ApplyURI(conn_string)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return mongo.Connect(ctx, clOptions)
}

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("error loading env file: %v\n", err)
	}
	client, err := dbClient()
	if err != nil {
		panic(err)
	}
	db_name := os.Getenv("URL_SHORTENER_DB_NAME")
	col_name := os.Getenv("URL_SHORTENER_COLLECTION_NAME")

	URL_COLLECTIONS = client.Database(db_name).Collection(col_name)
	log.Println("DATABASE IS READY FOR OPERATIONS")
}
