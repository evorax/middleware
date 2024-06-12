package main

import (
	middle "github.com/evorax/middleware"
)

func main() {
	e := middle.New()

	e.GET("/", func(ctx *middle.Context) {
		ctx.JSON(200, map[string]string{
			"key": "value",
		})
	})

	e.Listen(":8080")
}
