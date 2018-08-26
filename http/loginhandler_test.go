package http

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"strings"
	"fmt"
	."mygoapp/asserts"
)

func TestLoginLoadPage(t *testing.T) {

	// given
	req, err := http.NewRequest(http.MethodGet, "/",nil)
	if err != nil{
		t.Fatal(err.Error())
	}

	w := httptest.NewRecorder()
	handler := NewLoginHandler()

	// when
	handler.HandleLogin(w, req)

	// then
	b, err := ioutil.ReadAll(w.Body)
	fmt.Printf(string(b)+"hhhhhh"+"body")
	Equals(t, http.StatusOK, w.Code)
	Equals(t, true, strings.Contains(string(b), "Hello World"))
}


