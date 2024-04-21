package main

import (
	//"fmt"
	"net/http"
	"github.com/Blthrst/go-todos/handlers"
	"github.com/Blthrst/go-todos/model"
)

func main() {

	model.InitSecrets()

	http.HandleFunc("/users", handlers.UsersHandler)

    http.ListenAndServe("localhost:4545", nil)
}