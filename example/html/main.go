package main

func main() {
	e := middleware.New()

	e.GET("/", func(ctx *middleware.Context) {
		ctx.HTML(200, "<h1>hello</h1>")
	})

	e.Listen(":8080")
}
