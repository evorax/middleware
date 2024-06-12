package main

import (
	"fmt"

	"github.com/evorax/middleware"
)

func main() {
	e := middleware.New()

	e.Route("/", func(ctx *middleware.Context) {
		fmt.Println(ctx.Method())
	})

	e.Listen(":8080")
}
