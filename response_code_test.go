package belajar_go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "name is empty")
	} else {
		fmt.Fprintf(writer, "Hello %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest("GET", ApiHome, nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println("statusCode", response.StatusCode) //400
	fmt.Println("status", response.Status) //400 Bad Request
	fmt.Println(string(body))
}
func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest("GET", ApiHome+"/?name=Aku", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println("statusCode", response.StatusCode) //200
	fmt.Println("status", response.Status) //200 OK
	fmt.Println(string(body))
}
