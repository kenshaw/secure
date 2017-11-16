// examples/goji/main.go
package main

import (
	"net/http"

	"github.com/kenshaw/secure"
	"goji.io"
	"goji.io/pat"
)

func main() {
	mux := goji.NewMux()
	mux.Use(secure.New(
		secure.FrameDeny(true),
	).Handler)

	mux.HandleFunc(pat.Get("/"), func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("X-Frame-Options header is now `DENY`."))
	})

	http.ListenAndServe(":8080", mux)
}
