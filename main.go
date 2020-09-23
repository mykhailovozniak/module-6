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
	loadErr := godotenv.Load()

	if loadErr != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")

	db.Connect()

	mux := http.NewServeMux()

	finalHelloHandler := http.HandlerFunc(controllers.HelloController)
	finalMaterialsHandler := http.HandlerFunc(controllers.MaterialsController)
	finalPostHandler := http.HandlerFunc(controllers.PostController)

	mux.Handle("/hello", middlewares.CorrelationIDMiddleware(finalHelloHandler))
	mux.Handle("/materials", middlewares.CorrelationIDMiddleware(finalMaterialsHandler))
	mux.Handle("/post", middlewares.CorrelationIDMiddleware(finalPostHandler))

	err := http.ListenAndServe(":" + port, mux)

	if err != nil {
		log.Fatal("Error during start app")
	}
}
