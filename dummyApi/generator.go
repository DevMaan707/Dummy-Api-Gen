package dummyapi

import (
	"github.com/DevMaan707/dummy-api-gen/internal/generator"
	"github.com/DevMaan707/dummy-api-gen/internal/parser"
	"github.com/gin-gonic/gin"
)

func GenerateAPIs(router *gin.Engine, modelsPath string) error {
	parsedModels, err := parser.ParseModels(modelsPath)
	if err != nil {
		return err
	}
	return generator.SetupRoutes(router, parsedModels)
}
