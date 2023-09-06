package belajar_go_web

import (
	"net/http"
	"testing"
)

func ServerFile(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Query().Get("name") != "" {
		http.ServeFile(writer, request, "./resources/ok.html")
	} else {
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

func TestServerFileServer(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServerFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
