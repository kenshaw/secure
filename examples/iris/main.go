// examples/iris/main.go
package main

import (
	"github.com/kataras/iris"
	"github.com/kenshaw/secure"
)

func main() {
	secureMiddleware := &secure.Middleware{
		FrameDeny: true,
	}

	iris.UseFunc(func(c *iris.Context) {
		err := secureMiddleware.Process(c.ResponseWriter, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	})

	iris.Get("/home", func(c *iris.Context) {
		c.SendStatus(200, "X-Frame-Options header is now `DENY`.")
	})

	iris.Listen(":8080")
}
