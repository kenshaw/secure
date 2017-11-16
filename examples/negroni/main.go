// examples/negroni/main.go
package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/kenshaw/secure"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("X-Frame-Options header is now `DENY`."))
	})

	secureMiddleware := &secure.Middleware{
		FrameDeny: true,
	}

	n := negroni.Classic()
	n.Use(negroni.HandlerFunc(secureMiddleware.HandlerFuncWithNext))
	n.UseHandler(mux)

	n.Run("127.0.0.1:3000")
}
