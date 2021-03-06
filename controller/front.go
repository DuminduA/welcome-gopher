package controller

import (
	"encoding/json"
	"io"
	"net/http"
)

func RegisterControllers() {
	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)

}

func encodeToJson(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	err := enc.Encode(data)
	if err != nil {
		return
	}
}
