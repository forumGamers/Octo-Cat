package config

import (
	"context"
	"log"

	"os"
	"time"

	"github.com/forumGamers/Octo-Cat/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	DB *mongo.Database
)

func Connection() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DATABASE_URL")))
	errors.PanicIfError(err)
	errors.PanicIfError(client.Ping(ctx, readpref.Primary()))

	log.Println("connection success")
	DB = client.Database("Post")
}

func TestingConnection() *mongo.Database {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("DATABASE_URL")))
	errors.PanicIfError(err)

	return client.Database("Post_test")
}

func DisconnectConnection(client *mongo.Database) error {
	return client.Client().Disconnect(context.Background())
}
