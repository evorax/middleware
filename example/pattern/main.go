package main

import (
	"fmt"

	"github.com/evorax/middleware"
)

func main() {
	e := middleware.New()

	e.GET(`/foo/:id[^[a-zA-Z]+$]/:id2[^\d+$]`, func(ctx *middleware.Context) {
		ctx.WriteString(200, fmt.Sprintf("id:%s id2:%s", ctx.GetParam("id"), ctx.GetParam("id2")))
	})

	e.GET(`/abc/*/:id`, func(ctx *middleware.Context) {
		ctx.WriteString(200, ctx.GetParam("id"))
	})

	e.Listen(":8080")
}
