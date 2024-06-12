package middleware

import (
	"net/http"
	"strings"

	"github.com/evorax/middleware/internal"
)

func (e *Engine) AddRoute(method string, pattern string, handler HandlerFunc) {
	e.routes = append(e.routes, &route{
		method:  method,
		pattern: pattern,
		handler: handler,
	})
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	internal.Info(r.Method, r.URL.Path)
	for _, route := range e.routes {
		/*if route.method == r.Method {
			err := Method(route, w, r)
			if err != nil {
				internal.Error(err)
			}
		} else {
			err := Method(route, w, r)
			if err != nil {
				internal.Error(err)
			}
		}*/
		err := Method(route, w, r)
		if err != nil {
			internal.Error(err)
		}
	}

	if e.staticDir != "" {
		dir := strings.ReplaceAll(e.staticDir, ".", "")

		if !strings.HasSuffix(dir, "/") {
			dir += "/"
		}
		http.StripPrefix(strings.ReplaceAll(e.staticDir, ".", ""), http.FileServer(http.Dir(e.staticDir))).ServeHTTP(w, r)
	} else {
		http.NotFound(w, r)
	}
}

func Method(route *route, w http.ResponseWriter, r *http.Request) error {
	params, err := internal.MatchPath(route.pattern, r.URL.Path)
	if err != nil {
		return err
	}
	ctx := &Context{
		Writer:  w,
		Request: r,
		Params:  params,
	}
	route.handler(ctx)
	return nil
}
