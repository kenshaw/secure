// examples/std/main.go
package main

import (
	"net/http"

	"github.com/kenshaw/secure"
)

var myHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
})

func main() {
	secureMiddleware := &secure.Middleware{
		AllowedHosts:             []string{"example.com", "ssl.example.com"},
		HostsProxyHeaders:        []string{"X-Forwarded-Host"},
		SSLRedirect:              true,
		SSLHost:                  "ssl.example.com",
		SSLForwardedProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
		STSSeconds:               315360000,
		STSIncludeSubdomains:     true,
		STSPreload:               true,
		FrameDeny:                true,
		ContentTypeNosniff:       true,
		BrowserXSSFilter:         true,
		ContentSecurityPolicy:    "default-src 'self'",
	}

	app := secureMiddleware.Handler(myHandler)
	http.ListenAndServe("127.0.0.1:3000", app)
}
