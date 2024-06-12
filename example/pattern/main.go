package main

import (
	"fmt"

	middle "github.com/evorax/middleware/pkg"
)

func main() {
	e := middle.New()

	e.GET(`/foo/:id[^[a-zA-Z]+$]/:id2[^\d+$]`, func(ctx *middle.Context) {
		ctx.Write(200, fmt.Sprintf("id:%s id2:%s", ctx.GetParam("id"), ctx.GetParam("id2")))
	})

	e.GET(`/abc/*/:id`, func(ctx *middle.Context) {
		ctx.Write(200, ctx.GetParam("id"))
	})

	e.Listen(":8080")
}
