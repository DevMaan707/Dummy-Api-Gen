package dummyapi

import (
    "github.com/gin-gonic/gin"
    "dummy-api-generator/internal/generator"
    "dummy-api-generator/internal/parser"
)


func GenerateAPIs(router *gin.Engine, modelsPath string) error {
    parsedModels, err := parser.ParseModels(modelsPath)
    if err != nil {
        return err
    }
    return generator.SetupRoutes(router, parsedModels)
}
