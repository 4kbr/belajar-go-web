package belajar_go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	// server := http.Server{
	// 	Addr: "localhost:8020",
	// }
	// err := server.ListenAndServe()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Run in localhost")

	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		_, err := fmt.Fprint(writer, "holla wordo")
		if err != nil {
			panic(err)
		}
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello word")
	})
	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello Hi")
	})
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Ini images/")
	})
	mux.HandleFunc("/images/test", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Ini images/test")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, request.Method)
		fmt.Fprint(writer, request.RequestURI)
	}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
