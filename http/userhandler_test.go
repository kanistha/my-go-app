package http_test

import (
	"testing"
	gohttp "net/http"
	"net/http/httptest"
	"io/ioutil"
	"encoding/json"

	"mygoapp/http"
	"mygoapp/mock"
	."mygoapp/asserts"
)

func TestGetUsers(t *testing.T) {

	// given
	req, err := gohttp.NewRequest(gohttp.MethodGet, "/users?id=1", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	res := httptest.NewRecorder()
	repo := mock.NewRepositoryMock()
	handler := http.NewUserHandler(repo)

	// when
	handler.GetUsers(res, req)

	// then
	body, _ := ioutil.ReadAll(res.Body)

	var m map[string]string
	json.Unmarshal(body, &m)

	Equals(t, gohttp.StatusOK, res.Code)
	Equals(t, true, m["id"] == "1")
	Equals(t, true, "Foo Boo" == m["name"])
	Equals(t, true, "foo.boo@mail.com" == m["email"])

}


func TestGetUsers_withoutId(t *testing.T) {

	// given
	req, err := gohttp.NewRequest(gohttp.MethodGet, "/users", nil)
	if err != nil {
		t.Fatal(err.Error())
	}

	res := httptest.NewRecorder()
	repo := mock.NewRepositoryMock()
	handler := http.NewUserHandler(repo)

	// when
	handler.GetUsers(res, req)

	// then
	body, _ := ioutil.ReadAll(res.Body)

	var m map[string]string
	json.Unmarshal(body, &m)

	Equals(t, gohttp.StatusBadRequest, res.Code)
	Equals(t, true, m["error"] == "INVALID_REQUEST")
}

