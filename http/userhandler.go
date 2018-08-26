package http

import (
	gohttp "net/http"
	"mygoapp"
	"encoding/json"
)
const Error = "{\"error\":\"INVALID_REQUEST\"}"

type UserHandler struct {
	repository mygoapp.Repository
}

func NewUserHandler(r mygoapp.Repository) *UserHandler {
	return &UserHandler{repository: r}
}

func (u *UserHandler) GetUsers(w gohttp.ResponseWriter, r *gohttp.Request) {

	id := r.URL.Query().Get("id")
	if id == "" {
		gohttp.Error(w,Error, gohttp.StatusBadRequest)
		return
	}
	response := u.repository.Get(id)

	b, _ := json.Marshal(response)
	w.WriteHeader(gohttp.StatusOK)
	w.Write(b)

}
