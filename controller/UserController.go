package controller

import (
	"net/http"
	"regexp"
	"welcome-gopher/models"
)

type UserController struct {
	userIDPattern *regexp.Regexp
}

func (uc UserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello from user controller"))
	if err != nil {
		return
	}
}

func (uc UserController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeToJson(models.GetUsers(), w)
}

func newUserController() *UserController {
	return &UserController{
		userIDPattern: regexp.MustCompile("^/users/$"),
	}
}
