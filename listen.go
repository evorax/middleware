package middleware

import "net/http"

func (e *Engine) Listen(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ListenTLS(addr string, certFile string, keyFile string) error {
	return http.ListenAndServeTLS(addr, certFile, keyFile, e)
}
