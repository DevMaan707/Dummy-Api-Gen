package main

import (
	"log"

	dummyapi "github.com/DevMaan707/dummy-api-gen/dummyApi"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	modelsPath := "./examples/models"

	err := dummyapi.GenerateAPIs(router, modelsPath)
	if err != nil {
		log.Fatalf("Error generating APIs: %v", err)
	}

	log.Println("Server is running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
