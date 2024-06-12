package main

import (
	"net/http"

	"github.com/evorax/middleware"
)

func main() {
	e := middleware.New()

	e.GET("/", func(ctx *middleware.Context) {
		ctx.JSON(200, map[string]string{
			"key": "value",
		})
	})

	http.ListenAndServe(":8080", e)
}
