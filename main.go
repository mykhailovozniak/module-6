package main

import (
	"github.com/joho/godotenv"
	"log"
	"module-6/controllers"
	"module-6/db"
	"module-6/middlewares"
	"net/http"
	"os"
)

func main()  {
	log.Println("START", os.Getenv("MONGO_URI"))
	loadErr := godotenv.Load()

	if loadErr != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	db.Connect()

	mux := http.NewServeMux()

	finalHelloHandler := http.HandlerFunc(controllers.HelloController)
	finalMaterialsHandler := http.HandlerFunc(controllers.MaterialsController)

	mux.Handle("/hello", middlewares.CorrelationIDMiddleware(finalHelloHandler))
	mux.Handle("/materials", middlewares.CorrelationIDMiddleware(finalMaterialsHandler))


	err := http.ListenAndServe(":" + port, mux)

	if err != nil {
		log.Fatal("Error during start app")
	}
}
