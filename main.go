package main

import (
	//"fmt"
	"net/http"
	"github.com/Blthrst/go-todos/handlers"
	"github.com/Blthrst/go-todos/model"
)

func main() {

	//create db and tables automatically

	model.InitSecrets()

	http.HandleFunc("/users/", handlers.GetAllUsers)
	http.HandleFunc("/users/new", handlers.CreateUsers)
	http.HandleFunc("/users/delete", handlers.DeleteUsers)
	http.HandleFunc("/users/update", handlers.UpdateUser)
	http.HandleFunc("/users/one", handlers.GetOneUser)

    http.ListenAndServe("localhost:4545", nil)
}