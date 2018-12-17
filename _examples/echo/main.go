// examples/echo/main.go
package main

import (
	"net/http"

	"github.com/kenshaw/secure"
	"github.com/labstack/echo"
)

func main() {
	secureMiddleware := &secure.Middleware{
		FrameDeny: true,
	}

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "X-Frame-Options header is now `DENY`.")
	})

	e.Use(echo.WrapMiddleware(secureMiddleware.Handler))
	e.Logger.Fatal(e.Start("127.0.0.1:3000"))
}
