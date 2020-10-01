package models

import (
	"context"
	"log"
)

type Material struct {
	Name string `json:"Name"`
}

type Cursor interface {
	Next(ctx context.Context) bool
	Decode(val interface{}) (err error)
	Err() (err error)
}

func MapMaterials(cur Cursor) []*Material {
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
