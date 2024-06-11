package main

import (
	"fmt"

	"github.com/evorax/middleware/middle"
)

func main() {
	e := middle.New()

	e.GET("/{id}", func(ctx *middle.Context) {
		id := ctx.GetParam("id")
		ctx.HTML(200, fmt.Sprintf("welcome to %s!", id))
	})

	e.Listen(":8080")
}
