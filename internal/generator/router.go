package generator

import "github.com/gin-gonic/gin"

func InitializeRouter() *gin.Engine {
    router := gin.Default()
    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    return router
}
