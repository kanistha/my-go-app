package http

import (
	"mygoapp"
	"html/template"
	"net/http"
)

type Login struct {
	Loginer mygoapp.LoginService
	template *template.Template
}

func NewLoginHandler() *Login{
	return &Login{
		template: template.Must(template.ParseGlob("../template/*")),
	}
}

func (l *Login) HandleLogin(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case http.MethodGet:
		l.template.ExecuteTemplate(w, "loginform", nil)
	case http.MethodPost:
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
