package adapters

import (
	"bytes"
	"io"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

type FiberRouter struct {
	app *fiber.App
}

func NewFiberRouter(app *fiber.App) *FiberRouter {
	return &FiberRouter{app: app}
}

func (r *FiberRouter) Group(prefix string) RouterGroup {
	return &FiberRouterGroup{group: r.app.Group(prefix)}
}

func (r *FiberRouter) Run(address string) error {
	return r.app.Listen(address)
}

type FiberRouterGroup struct {
	group fiber.Router
}

func (g *FiberRouterGroup) GET(path string, handler func(http.ResponseWriter, *http.Request)) {
	g.group.Get(path, func(c *fiber.Ctx) error {
		req, err := convertFiberRequest(c)
		if err != nil {
			c.Status(http.StatusInternalServerError).SendString("Failed to convert request")
			return nil
		}
		w := &fiberResponseWriter{ctx: c}
		handler(w, req)
		return nil
	})
}

func (g *FiberRouterGroup) POST(path string, handler func(http.ResponseWriter, *http.Request)) {
	g.group.Post(path, func(c *fiber.Ctx) error {
		req, err := convertFiberRequest(c)
		if err != nil {
			c.Status(http.StatusInternalServerError).SendString("Failed to convert request")
			return nil
		}
		w := &fiberResponseWriter{ctx: c}
		handler(w, req)
		return nil
	})
}

func convertFiberRequest(ctx *fiber.Ctx) (*http.Request, error) {
	req := &http.Request{}
	u, err := url.ParseRequestURI(ctx.OriginalURL())
	if err != nil {
		return nil, err
	}
	req.URL = u
	req.Method = ctx.Method()
	req.Header = http.Header{}
	ctx.Request().Header.VisitAll(func(key, value []byte) {
		req.Header.Add(string(key), string(value))
	})
	req.Body = io.NopCloser(bytes.NewReader(ctx.Body()))
	return req, nil
}

type fiberResponseWriter struct {
	ctx *fiber.Ctx
}

func (w *fiberResponseWriter) Header() http.Header {
	headers := http.Header{}
	w.ctx.Response().Header.VisitAll(func(key, value []byte) {
		headers.Set(string(key), string(value))
	})
	return headers
}

func (w *fiberResponseWriter) Write(data []byte) (int, error) {
	return w.ctx.Write(data)
}

func (w *fiberResponseWriter) WriteHeader(statusCode int) {
	w.ctx.Status(statusCode)
}
