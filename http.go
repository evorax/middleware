package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/evorax/middleware/internal"
)

func (ctx *Context) GetParam(key string) string {
	value, ok := ctx.Params[key]
	if !ok {
		return fmt.Sprintf("param '%s' not found", key)
	}
	return value
}

func (ctx *Context) JSON(status int, obj any, headers ...map[string]string) error {
	return internal.WriteJSON(ctx.Writer, status, obj, headers...)
}

func (ctx *Context) HTML(status int, obj string, headers ...map[string]string) error {
	return internal.WriteHTML(ctx.Writer, status, obj, headers...)
}

func (ctx *Context) Write(status int, obj []byte, headers ...map[string]string) error {
	return internal.WriteTo(ctx.Writer, status, obj, headers...)
}

func (ctx *Context) WriteString(status int, obj string, headers ...map[string]string) error {
	return internal.WriteTo(ctx.Writer, status, []byte(obj), headers...)
}

func (ctx *Context) SetHeader(key, value string) {
	ctx.Writer.Header().Set(key, value)
}

func (ctx *Context) SetCookie(name, value, path, domain string, date time.Duration, options ...Cookie) {
	cookie := &http.Cookie{
		Name:  name,
		Value: value,

		Path:    path,
		Domain:  domain,
		Expires: time.Now().Add(date),
	}

	if cookie.MaxAge != 0 {
		cookie.MaxAge = options[0].MaxAge
	}
	if cookie.Secure {
		cookie.Secure = options[0].Secure
	}
	if cookie.Raw != "" {
		cookie.Raw = options[0].Raw
	}
	if len(cookie.Unparsed) > 0 {
		cookie.Unparsed = options[0].Unparsed
	}
	if cookie.HttpOnly {
		cookie.HttpOnly = options[0].HttpOnly
	}

	http.SetCookie(ctx.Writer, cookie)
}
