package middleware

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.AddRoute("GET", pattern, handler)
}

func (e *Engine) HEAD(pattern string, handler HandlerFunc) {
	e.AddRoute("HEAD", pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.AddRoute("POST", pattern, handler)
}

func (e *Engine) PUT(pattern string, handler HandlerFunc) {
	e.AddRoute("PUT", pattern, handler)
}

func (e *Engine) DELETE(pattern string, handler HandlerFunc) {
	e.AddRoute("DELETE", pattern, handler)
}

func (e *Engine) CONNECT(pattern string, handler HandlerFunc) {
	e.AddRoute("CONNECT", pattern, handler)
}

func (e *Engine) OPTIONS(pattern string, handler HandlerFunc) {
	e.AddRoute("OPTIONS", pattern, handler)
}

func (e *Engine) TRACE(pattern string, handler HandlerFunc) {
	e.AddRoute("TRACE", pattern, handler)
}

func (e *Engine) PATCH(pattern string, handler HandlerFunc) {
	e.AddRoute("PATCH", pattern, handler)
}

func (e *Engine) Other(method string, pattern string, handler HandlerFunc) {
	e.AddRoute(method, pattern, handler)
}

func (e *Engine) Route(pattern string, handler HandlerFunc) {
	e.routes = append(e.routes, &route{
		method:  "",
		pattern: pattern,
		handler: handler,
	})
}
