package adapters

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	engine *gin.Engine
}

func NewGinRouter(engine *gin.Engine) *GinRouter {
	return &GinRouter{engine: engine}
}

func (r *GinRouter) Group(prefix string) RouterGroup {
	return &GinRouterGroup{group: r.engine.Group(prefix)}
}

func (r *GinRouter) Run(address string) error {
	return r.engine.Run(address)
}

type GinRouterGroup struct {
	group *gin.RouterGroup
}

func (g *GinRouterGroup) GET(path string, handler func(http.ResponseWriter, *http.Request)) {
	g.group.GET(path, gin.WrapF(handler))
}

func (g *GinRouterGroup) POST(path string, handler func(http.ResponseWriter, *http.Request)) {
	g.group.POST(path, gin.WrapF(handler))
}
