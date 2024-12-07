package main

import (
	"log"
	"net/http"

	"github.com/DevMaan707/dummy-api-gen/adapters"
	"github.com/DevMaan707/dummy-api-gen/api"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	router := adapters.NewGinRouter(app)

	models, err := api.ParseModels("./models")
	if err != nil {
		log.Fatalf("Error parsing models: %v", err)
	}

	err = api.GenerateAPIs(router, models)
	if err != nil {
		log.Fatalf("Error generating APIs: %v", err)
	}

	router.Group("/custom").GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	log.Println("Server running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
