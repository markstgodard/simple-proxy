package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const (
	ROUTES       = "ROUTES"
	DEFAULT_PORT = "3000"
)

type Routes []Route
type Route struct {
	Path string
	Host string
}

func loadConfig() (Routes, error) {
	data := os.Getenv(ROUTES)
	routes := Routes{}
	err := json.Unmarshal([]byte(data), &routes)
	if err != nil {
		fmt.Println("Error parsing routes:", err)
	}
	return routes, err
}

func main() {

	log.SetOutput(os.Stdout)

	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	routes, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	for _, r := range routes {
		fmt.Printf("mapping path [%s] to host [%s]\n", r.Path, r.Host)
		target, err := url.Parse(r.Host)
		if err != nil {
			log.Fatal(err)
		}
		http.Handle(r.Path+"/", httputil.NewSingleHostReverseProxy(target))
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
