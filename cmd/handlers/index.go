package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func IndexHandler(writer http.ResponseWriter, request *http.Request) {
	readFile, err := template.ParseFiles("index.html")

	if err != nil {
		log.Fatalf("Cannot find file")
	}

	index := template.Must(readFile, err)

	if err != nil {
		log.Fatalf("Cannot open template")
	}

	index.Execute(writer, nil)
}
