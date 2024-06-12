# Middleware

under development...

Libraries we plan to use for eptrone.com</br>
Basic code for
```go
package main

import (
	"github.com/evorax/middleware"
)

func main() {
	e := middleware.New()

	e.GET("/", func(ctx *middleware.Context) {
		ctx.JSON(200, map[string]string{
			"key": "value",
		})
	})

	e.Listen(":8080")
}
```
</br>
HTML code for

```go
package main

import "github.com/evorax/middleware"

func main() {
	e := middleware.New()

	e.GET("/", func(ctx *middleware.Context) {
		ctx.HTML(200, "<h1>hello</h1>")
	})

	e.Listen(":8080")
}
```