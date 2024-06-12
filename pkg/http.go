package middle

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (ctx *Context) GetParam(key string) string {
	value, ok := ctx.Params[key]
	if !ok {
		return fmt.Sprintf("param '%s' not found", key)
	}
	return value
}

func (ctx *Context) JSON(status int, obj any, options ...Headers) error {
	if len(options) != 0 {
		for key, value := range options[0].Header {
			ctx.Writer.Header().Add(key, value)
		}
	}

	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(status)

	parse, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	ctx.Writer.Write(parse)
	return nil
}

func (ctx *Context) HTML(status int, obj string, options ...Headers) {
	if len(options) != 0 {
		for key, value := range options[0].Header {
			ctx.Writer.Header().Add(key, value)
		}
	}
	ctx.Writer.WriteHeader(status)
	ctx.Writer.Write([]byte(obj))
}

func (ctx *Context) Write(status int, obj string, options ...Headers) {
	if len(options) != 0 {
		for key, value := range options[0].Header {
			ctx.Writer.Header().Add(key, value)
		}
	}
	ctx.Writer.WriteHeader(status)
	ctx.Writer.Write([]byte(obj))
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
