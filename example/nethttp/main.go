package main

import (
	"net/http"

	middle "github.com/evorax/middleware/pkg"
)

func main() {
	e := middle.New()

	e.GET("/", func(ctx *middle.Context) {
		ctx.JSON(200, map[string]string{
			"key": "value",
		})
	})

	http.ListenAndServe(":8080", e)
}
