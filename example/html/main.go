package main

import (
	"github.com/evorax/middleware/middle"
)

func main() {
	e := middle.New()

	e.GET("/", func(ctx *middle.Context) {
		ctx.HTML(200, "<h1>hello</h1>")
	})

	e.Listen(":8080")
}
