package middleware

import (
	"crypto/tls"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
)

func (ctx *Context) Method() string {
	return ctx.Request.Method
}

func (ctx *Context) Query() url.Values {
	return ctx.Request.URL.Query()
}

func (ctx *Context) Host() string {
	return ctx.Request.Host
}

func (ctx *Context) URL() *url.URL {
	return ctx.Request.URL
}

func (ctx *Context) Proto() string {
	return ctx.Request.Proto
}

func (ctx *Context) ProtoMajor() int {
	return ctx.Request.ProtoMajor
}

func (ctx *Context) ProtoMinor() int {
	return ctx.Request.ProtoMajor
}

func (ctx *Context) Body() io.ReadCloser {
	return ctx.Request.Body
}

func (ctx *Context) GetBody() func() (io.ReadCloser, error) {
	return ctx.Request.GetBody
}

func (ctx *Context) ContentLength() int64 {
	return ctx.Request.ContentLength
}

func (ctx *Context) TransferEncoding() []string {
	return ctx.Request.TransferEncoding
}

func (ctx *Context) Close() bool {
	return ctx.Request.Close
}

func (ctx *Context) Form() url.Values {
	return ctx.Request.Form
}

func (ctx *Context) PostForm() url.Values {
	return ctx.Request.PostForm
}

func (ctx *Context) MultipartForm() *multipart.Form {
	return ctx.Request.MultipartForm
}

func (ctx *Context) RemoteAddr() string {
	return ctx.Request.RemoteAddr
}

func (ctx *Context) RequestURI() string {
	return ctx.Request.RequestURI
}

func (ctx *Context) TLS() *tls.ConnectionState {
	return ctx.Request.TLS
}

func (ctx *Context) Cancel() <-chan struct{} {
	return ctx.Request.Cancel
}

func (ctx *Context) Response() *http.Response {
	return ctx.Request.Response
}
