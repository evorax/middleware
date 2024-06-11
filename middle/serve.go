package middle

import (
	"net/http"
	"strings"
)

func (e *Engine) AddRoute(method string, pattern string, handler HandlerFunc) {
	e.routes = append(e.routes, &route{
		method:  method,
		pattern: pattern,
		handler: handler,
	})
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Info(r.Method, r.URL.Path)
	for _, route := range e.routes {
		if route.method == r.Method {
			params, err := MatchPath(route.pattern, r.URL.Path)
			if err != nil {
				Error(err)
				return
			}
			ctx := &Context{
				Writer:  w,
				Request: r,
				Params:  params,
			}
			route.handler(ctx)
			return
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
