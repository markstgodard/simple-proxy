package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	target, err := url.Parse("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	appEnv := cfenv.Current()

	http.Handle("/products/", httputil.NewSingleHostReverseProxy(target))
	log.Fatal(http.ListenAndServe(":3000", nil))
}
