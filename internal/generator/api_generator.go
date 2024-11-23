package generator

import (
    "github.com/gin-gonic/gin"
    "net/http"
)


func SetupRoutes(router *gin.Engine, models []string) error {
    group := router.Group("/test")
    LogInfo("Setting up dummy routes...")
for _, model := range models {
    route := "/" + model
    LogInfo("Creating GET API for " + route)
    group.GET(route, func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Dummy response for " + model,
        })
    })
}

    return nil
}
