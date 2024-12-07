package adapters

import "net/http"

type Router interface {
	Group(prefix string) RouterGroup
	Run(addr string) error
}

type RouterGroup interface {
	GET(path string, handler func(http.ResponseWriter, *http.Request))
	POST(path string, handler func(http.ResponseWriter, *http.Request))
}
