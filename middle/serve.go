package middle

import (
	"net/http"
	"regexp"
	"strings"
)

func (e *Engine) AddRoute(method string, pattern string, handler HandlerFunc) {
	regexPattern := "^" + pattern
	regexPattern = strings.ReplaceAll(regexPattern, "{", "(?P<")
	regexPattern = strings.ReplaceAll(regexPattern, "}", ">[^/]+)") + "$"
	re := regexp.MustCompile(regexPattern)

	e.routes = append(e.routes, &route{
		method:  method,
		pattern: re,
		handler: handler,
	})
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Info(r.Method, r.URL.Path)
	for _, route := range e.routes {
		if route.method == r.Method {
			matches := route.pattern.FindStringSubmatch(r.URL.Path)
			if matches != nil {
				params := make(map[string]string)
				for i, name := range route.pattern.SubexpNames() {
					if i != 0 && name != "" {
						params[name] = matches[i]
					}
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
