// examples/redirect/main.go
package main

import (
	"log"
	"net/http"

	"github.com/kenshaw/secure"
)

var myHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
})

func main() {
	secureMiddleware := &secure.Middleware{
		SSLRedirect: true,

		// This is optional in production. The default behavior is to just
		// redirect the request to the HTTPS protocol. Example:
		// http://github.com/some_page would be redirected to
		// https://github.com/some_page.
		SSLHost: "localhost:8443",
	}

	app := secureMiddleware.Handler(myHandler)

	// HTTP
	go func() {
		log.Fatal(http.ListenAndServe(":8080", app))
	}()

	// HTTPS
	// To generate a development cert and key, run the following from your *nix terminal:
	// go run $GOROOT/src/crypto/tls/generate_cert.go --host="localhost"
	log.Fatal(http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", app))
}
