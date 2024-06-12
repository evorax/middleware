package main

import (
	middle "github.com/evorax/middleware/pkg"
)

func main() {
	e := middle.New()

	e.GET("/", func(ctx *middle.Context) {
		ctx.HTML(200, "<h1>hello</h1>")
	})

	e.Listen(":8080")
}
