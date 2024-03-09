package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	defaultPort = ":3000"
)

func main() {
	pathToUrl := map[string]string{
		"/google":   "https://www.google.com",
		"/facebook": "https://www.facebook.com",
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler(pathToUrl))
	fmt.Printf("Server is running on port %s\n", defaultPort)
	if err := http.ListenAndServe(defaultPort, MapHandler(pathToUrl, mux)); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if url, ok := pathsToUrls[r.URL.Path]; ok {
			http.Redirect(w, r, url, http.StatusSeeOther)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func HomeHandler(m map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "text/html")
		homePage := "<h1>Redirects</h1><ul>"
		for path, url := range m {
			homePage += fmt.Sprintf("<li>%s : %s</li>", path, url)
		}
		homePage += "</ul>"

		w.Write([]byte(homePage))
	}
}
