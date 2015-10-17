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
	Path        string
	Host        string
	StripPrefix bool
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
		// fmt.Printf("mapping path [%s] to host [%s]\n", r.Path, r.Host)
		target, err := url.Parse(r.Host)
		if err != nil {
			log.Fatal(err)
		}

		path := r.Path
		if path != "/" {
			path = path + "/"
		}

		if r.StripPrefix {
			fmt.Printf("(strip) mapping %s to %s\n", path, target)
			http.Handle(path, http.StripPrefix(path, httputil.NewSingleHostReverseProxy(target)))
		} else {
			fmt.Printf("mapping %s to %s\n", path, target)
			http.Handle(path, httputil.NewSingleHostReverseProxy(target))
		}
	}

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
