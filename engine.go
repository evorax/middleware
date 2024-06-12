package middleware

import (
	"net/http"

	"github.com/evorax/middleware/internal"
)

type Context struct {
	Writer  http.ResponseWriter
	Request *http.Request
	Params  internal.Params
}

type HandlerFunc func(*Context)

type route struct {
	method  string
	pattern string
	handler HandlerFunc
}

type Engine struct {
	routes    []*route
	staticDir string
}

type Cookie struct {
	RawExpires string

	MaxAge   int
	Secure   bool
	SameSite SameSite
	Raw      string
	Unparsed []string
	HttpOnly bool
}

type SameSite int

type Headers map[string]string

func New(staticDir ...string) *Engine {
	if len(staticDir) >= 1 {
		return &Engine{routes: []*route{}, staticDir: staticDir[0]}
	}
	return &Engine{routes: []*route{}, staticDir: ""}
}
