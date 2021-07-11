package main

import (
	"fmt"
	"net/http"
	"welcome-gopher/controller"
	"welcome-gopher/models"
)

func main() {
	TestVarSyntax()
	fmt.Println("Hello gophers")

	user := models.User{
		ID:        1,
		FirstName: "Red",
		LastName:  "blue",
	}
	fmt.Println(user)

	controller.RegisterControllers()
	http.ListenAndServe(":3030", nil)

}
