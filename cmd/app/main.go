package main

import (
	"buraksrc/htmx/cmd/handlers"
	"fmt"
	"log"
	"net/http"
)

const PORT = ":8080"

var SERVER *http.ServeMux

func main() {
	SERVER = http.NewServeMux()

	SERVER.Handle("/", http.HandlerFunc(handlers.IndexHandler))
	SERVER.Handle("/public/css/", http.HandlerFunc(handlers.CssHandler))
	SERVER.Handle("/public/font/", http.HandlerFunc(handlers.FontHandler))

	srv := http.Server{
		Addr:    PORT,
		Handler: SERVER,
	}

	fmt.Printf("Server is running on %s\n", PORT)

	log.Fatal(srv.ListenAndServe())
}
