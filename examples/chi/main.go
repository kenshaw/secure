// examples/chi/main.go
package main

import (
	"net/http"

	"github.com/kenshaw/secure"
	"github.com/pressly/chi"
)

func main() {
	secureMiddleware := &secure.Middleware{
		FrameDeny: true,
	}

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("X-Frame-Options header is now `DENY`."))
	})
	r.Use(secureMiddleware.Handler)

	http.ListenAndServe("127.0.0.1:3000", r)
}
