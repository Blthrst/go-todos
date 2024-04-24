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

		err := json.NewDecoder(req.Body).Decode(&todo)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		err = model.CreateTodo(todo) 

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		data, _ := json.Marshal(ServerCreatedMessage)

		w.Write(data)
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
	}
}