package generator

import (
	"net/http"

	"github.com/DevMaan707/dummy-api-gen/internal/shared"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, models []shared.ModelData) error {
	group := router.Group("/test")
	LogInfo("Setting up dummy routes...")

	for _, model := range models {
		route := "/" + model.Name

		LogInfo("Creating GET API for " + route)
		getResponse := generateResponse(model.ResponseFields)
		group.GET(route, func(c *gin.Context) {
			c.JSON(http.StatusOK, getResponse)
		})

		if len(model.RequestFields) > 0 {
			LogInfo("Creating POST API for " + route)
			postResponse := generateResponse(model.ResponseFields)
			postValidator := model.RequestFields

			group.POST(route, func(c *gin.Context) {
				var requestBody map[string]interface{}
				if err := c.ShouldBindJSON(&requestBody); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
					return
				}

				for fieldName, fieldType := range postValidator {
					value, exists := requestBody[fieldName]
					if !exists {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Missing field: " + fieldName})
						return
					}

					if !validateFieldType(value, fieldType) {
						c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid type for field: " + fieldName})
						return
					}
				}

				c.JSON(http.StatusOK, postResponse)
			})
		}
	}

	return nil
}

func generateResponse(fields map[string]string) map[string]interface{} {
	response := make(map[string]interface{})
	for fieldName, fieldType := range fields {
		switch fieldType {
		case "int":
			response[fieldName] = 0
		case "string":
			response[fieldName] = "sample_text"
		default:
			response[fieldName] = nil
		}
	}
	return response
}

func validateFieldType(value interface{}, expectedType string) bool {
	switch expectedType {
	case "int":
		_, ok := value.(float64)
		return ok
	case "string":
		_, ok := value.(string)
		return ok
	default:
		return false
	}
}
