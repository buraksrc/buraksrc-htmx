package handlers

import (
	"net/http"
	"strings"
)

func CssHandler(writer http.ResponseWriter, request *http.Request) {
	css := http.FileServer(http.Dir("public/css"))

	if strings.HasSuffix(request.URL.Path, ".css") {
		writer.Header().Set("Content-Type", "text/css")
	}

	http.StripPrefix("/public/css", css).ServeHTTP(writer, request)
}

func FontHandler(writer http.ResponseWriter, request *http.Request) {
	filePath := request.URL.Path[1:]
	http.ServeFile(writer, request, filePath)
}
