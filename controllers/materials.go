package controllers

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"module-6/models"
	"net/http"
)

var materialsCollection *mongo.Collection

func MaterialsCollection(c *mongo.Database) {
	materialsCollection = c.Collection("materials")
}

func MaterialsController(res http.ResponseWriter, req *http.Request) {
	findOptions := options.Find()
	filter := bson.M{}
	cursor, materialErr := materialsCollection.Find(context.TODO(), filter, findOptions)

	if materialErr != nil {
		log.Fatal("Fatal Error during find materialsList")
	}

	materialsList := models.MapMaterials(cursor)
	jsonData, err := json.Marshal(materialsList)

	if err != nil {
		log.Println("Error during parsing data from draft")
		http.Error(res, err.Error(), http.StatusInternalServerError)

		return
	}

	res.Header().Set("Content-Type", "application/json")
	_, writeErr := res.Write(jsonData)

	if writeErr != nil {
		log.Println("Error during sending json")
		http.Error(res, writeErr.Error(), http.StatusInternalServerError)

		return
	}

	log.Print("Successfully a list of materialsList fetched from draft")
}
