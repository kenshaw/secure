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

	app := iris.New()
	app.Use(func(c iris.Context) {
		err := secureMiddleware.Process(c.ResponseWriter(), c.Request())

		// if there was an error, do not continue
		if err != nil {
			return
		}

		c.Next()
	})

	app.Get("/home", func(c iris.Context) {
		c.StatusCode(200)
		c.WriteString("X-Frame-Options header is now `DENY`.")
	})

	app.Run(iris.Addr(":8080"))
}
