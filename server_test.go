package belajar_go_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8020",
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	fmt.Println("Run in localhost")

}
