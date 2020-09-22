package models

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Material struct {
	Name string `json:"Name"`
}

func MapMaterials(cur *mongo.Cursor) []*Material {
	var materials []*Material

	for cur.Next(context.TODO()) {
		var material Material
		err := cur.Decode(&material)
		if err != nil {
			log.Fatal(err)
		}

		materials = append(materials, &material)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return materials
}
