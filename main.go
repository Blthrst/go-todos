package main

import (
	"net/http"
	"github.com/Blthrst/go-todos/handlers"
	"github.com/Blthrst/go-todos/model"
)

func main() {

	model.InitSecrets()
	model.PrepareDatabase()

	http.HandleFunc("/users/", handlers.GetAllUsers)
	http.HandleFunc("/users/create", handlers.CreateUsers)
	http.HandleFunc("/users/delete", handlers.DeleteUsers)
	http.HandleFunc("/users/update", handlers.UpdateUser)
	http.HandleFunc("/users/get", handlers.GetOneUser)

	http.HandleFunc("/todos/", handlers.GetTodos)
	http.HandleFunc("/todos/create", handlers.CreateTodo)
	http.HandleFunc("/todos/delete", handlers.DeleteTodos)
	http.HandleFunc("/todos/update", handlers.UpdateTodo)

    http.ListenAndServe("localhost:4545", nil)
}