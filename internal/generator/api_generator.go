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

		dynamicResponse := make(map[string]interface{})
		for fieldName, fieldType := range model.Fields {
			switch fieldType {
			case "int":
				dynamicResponse[fieldName] = 0
			case "string":
				dynamicResponse[fieldName] = "sample_text"
			default:
				dynamicResponse[fieldName] = nil
			}
		}

		group.GET(route, func(c *gin.Context) {
			c.JSON(http.StatusOK, dynamicResponse)
		})
	}

	return nil
}
