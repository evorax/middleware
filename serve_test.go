package middleware

import (
	"fmt"
	"testing"
)

func TestServe(t *testing.T) {
	e := New()
	e.AddRoute("GET", "/:id", func(ctx *Context) {
		id := ctx.GetParam("id")
		fmt.Println(id)
	})
}
