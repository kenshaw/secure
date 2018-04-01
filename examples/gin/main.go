// examples/gin/main.go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kenshaw/secure"
)

func main() {
	secureMiddleware := &secure.Middleware{
		FrameDeny: true,
	}
	secureFunc := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			err := secureMiddleware.Process(c.Writer, c.Request)

			// if there was an error, do not continue
			if err != nil {
				c.Abort()
				return
			}

			// avoid header rewrite if response is a redirection
			if status := c.Writer.Status(); status > 300 && status < 399 {
				c.Abort()
			}
		}
	}()

	router := gin.Default()
	router.Use(secureFunc)

	router.GET("/", func(c *gin.Context) {
		c.String(200, "X-Frame-Options header is now `DENY`.")
	})

	router.Run("127.0.0.1:3000")
}
