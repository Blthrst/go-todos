package handlers

import (
	"encoding/json"
	"github.com/Blthrst/go-todos/model"
	"log"
	"net/http"
	//"net/url"
	//"strconv"
)

func CreateTodo(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		todo := model.Todo{}

		decodeJSON(w, req, &todo)

		err := model.CreateTodo(todo)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		data, _ := json.Marshal(ServerCreatedMessage)

		w.Write(data)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func DeleteTodos(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		todosIds := []int{}

		err := json.NewDecoder(req.Body).Decode(&todosIds)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		err = model.DeleteTodos(todosIds)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		data, _ := json.Marshal(ServerCreatedMessage)

		w.Write(data)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func CompleteTodo(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		completeBody := model.CompleteTodoBody{}

		err := json.NewDecoder(req.Body).Decode(&completeBody)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		err = model.CompleteTodo(completeBody.Id)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		data, _ := json.Marshal(ServerCreatedMessage)

		w.Write(data)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func GetTodos(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		todos := []model.Todo{}

		err := json.NewDecoder(req.Body).Decode(&todos)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			return
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}