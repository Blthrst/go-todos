package handlers

import (
	"encoding/json"
	"github.com/Blthrst/go-todos/model"
	"log"
	"net/http"
)

func CreateTodo(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		todo := model.Todo{}

		todo = decodeJSON(w, req, todo)

		err := model.CreateTodo(todo)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		sendServerMessage(w, ServerCreatedMessage)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func DeleteTodos(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		delBody := model.DeleteTodosBody{}

		delBody = decodeJSON(w, req, delBody)

		err := model.DeleteTodos(delBody.TodoIds)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		sendServerMessage(w, ServerDeletedMessage)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func CompleteTodo(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		completeBody := model.CompleteTodoBody{}

		completeBody = decodeJSON(w, req, completeBody)

		err := model.CompleteTodo(completeBody.Id)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		sendServerMessage(w, ServerUpdatedMessage)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func GetTodos(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		todos, err := model.GetTodos()

		if err != nil {
			writeErrorAndLog(w, err)
		}

		data, _ := json.Marshal(todos)

		w.Write(data)

	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func UpdateTodo(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		todo := model.Todo{}

		todo = decodeJSON(w, req, todo)

		err := model.UpdateTodo(todo)

		if err != nil {
			writeErrorAndLog(w, err)
		}

		sendServerMessage(w, ServerUpdatedMessage)

	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}