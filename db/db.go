package db

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"module-6/controllers"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect()  {
	loadErr := godotenv.Load()
	if loadErr != nil {
		panic(loadErr)
	}

	mongoUri := os.Getenv("MONGO_URI")

	clientOptions := options.
		Client().
		ApplyURI(mongoUri)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	log.Println("Connected to MongoDB!")

	database := client.Database("test")
	controllers.MaterialsCollection(database)
}
