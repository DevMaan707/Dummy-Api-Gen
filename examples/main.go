package main

import (
	"log"
	"net/http"

	"github.com/DevMaan707/dummy-api-gen/adapters"
	"github.com/DevMaan707/dummy-api-gen/api"
	"github.com/gin-gonic/gin"
)

func main() {
	customEndpointConfigs := map[string]*api.EndpointConfig{
		"/Product": {
			StaticResponse: map[string]interface{}{
				"id":           1,
				"product_name": "Sample Product",
				"price":        99.99,
				"quantity":     100,
				"created_at":   "2024-12-07T12:00:00Z",
			},
		},

		"/User": {
			ConditionalLogic: func(request map[string]interface{}) map[string]interface{} {
				if userID, ok := request["user_id"].(string); ok {
					switch userID {
					case "0":
						return map[string]interface{}{
							"id":         0,
							"username":   "Guest",
							"email":      "guest@example.com",
							"created_at": "2024-12-07T12:00:00Z",
						}
					case "1":
						return map[string]interface{}{
							"id":         1,
							"username":   "UserOne",
							"email":      "userone@example.com",
							"created_at": "2024-12-06T12:00:00Z",
						}
					case "2":
						return map[string]interface{}{
							"id":         2,
							"username":   "UserTwo",
							"email":      "usertwo@example.com",
							"created_at": "2024-12-05T12:00:00Z",
						}
					case "3":
						return map[string]interface{}{
							"id":         3,
							"username":   "UserThree",
							"email":      "userthree@example.com",
							"created_at": "2024-12-04T12:00:00Z",
						}
					case "4":
						return map[string]interface{}{
							"id":         4,
							"username":   "UserFour",
							"email":      "userfour@example.com",
							"created_at": "2024-12-03T12:00:00Z",
						}
					default:
						return map[string]interface{}{
							"id":         -1,
							"username":   "Unknown",
							"email":      "unknown@example.com",
							"created_at": "2024-12-01T12:00:00Z",
						}
					}
				}
				return nil
			},
		},

		"/OrderDetails": {
			ConditionalLogic: func(request map[string]interface{}) map[string]interface{} {
				if orderID, ok := request["order_id"].(string); ok {
					switch orderID {
					case "0":
						return map[string]interface{}{
							"order_id": "0",
							"name":     "Sample Order Zero",
							"quantity": 0,
							"price":    0.0,
							"status":   "Cancelled",
						}
					case "1":
						return map[string]interface{}{
							"order_id": "1",
							"name":     "Order One",
							"quantity": 1,
							"price":    25.5,
							"status":   "Processing",
						}
					case "2":
						return map[string]interface{}{
							"order_id": "2",
							"name":     "Order Two",
							"quantity": 2,
							"price":    50.0,
							"status":   "Shipped",
						}
					case "3":
						return map[string]interface{}{
							"order_id": "3",
							"name":     "Order Three",
							"quantity": 3,
							"price":    75.5,
							"status":   "Delivered",
						}
					default:
						return map[string]interface{}{
							"order_id": "unknown",
							"name":     "Unknown Order",
							"quantity": -1,
							"price":    0.0,
							"status":   "Unknown",
						}
					}
				}
				return nil
			},
		},

		"/Order": {
			StaticResponse: map[string]interface{}{
				"order_id": "random-generated-id",
				"message":  "success",
			},
		},
	}
	app := gin.Default()
	router := adapters.NewGinRouter(app)

	models, err := api.ParseModels("./models")
	if err != nil {
		log.Fatalf("Error parsing models: %v", err)
	}

	err = api.GenerateAPIsWithConfig(router, models, customEndpointConfigs)
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
