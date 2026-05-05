package main

import (
	"net/http"

	"github.com/timu-ryan/urlshortener/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.PostLink)
	mux.HandleFunc("/{shortname}", handler.GetLink)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
