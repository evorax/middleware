package main

import (
	"fmt"

	"github.com/evorax/middleware/middle"
)

func main() {
	e := middle.New()

	e.GET(`/foo/:id[^[a-zA-Z]+$]/:id2[^\d+$]`, func(ctx *middle.Context) {
		fmt.Fprintf(ctx.Writer, fmt.Sprintf("id:%v id2:%v", ctx.Params["id"], ctx.Params["id2"]))
	})

	e.GET(`/abc/*/:id`, func(ctx *middle.Context) {
		fmt.Fprintf(ctx.Writer, fmt.Sprintf("id:%v", ctx.Params["id"]))
	})

	e.Listen(":8080")
}
