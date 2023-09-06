package belajar_go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writter http.ResponseWriter, request *http.Request) {

	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	//cara parsing manual:
	// firstName := request.PostForm.Get("first_name")
	// lastName := request.PostForm.Get("last_name")

	//dengan library langsung
	firstName := request.PostFormValue("first_name")
	lastName := request.PostFormValue("last_name")

	fmt.Fprintf(writter, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Aku&last_name=Kedua")
	request := httptest.NewRequest(http.MethodPost, ApiHome, requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
