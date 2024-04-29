package handlers

import (
	"encoding/json"
	"github.com/Blthrst/go-todos/model"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func GetAllUsers(w http.ResponseWriter, req *http.Request) {
	users, err := model.GetAllUsers()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	data, _ := json.Marshal(users)

	w.Write(data)
}

func GetOneUser(w http.ResponseWriter, req *http.Request) {

	userId, err := url.ParseQuery(req.URL.RawQuery)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	id, err := strconv.Atoi(userId["userId"][0])

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	u, err := model.GetUser(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	data, _ := json.Marshal(u)

	w.Write(data)

}

func CreateUsers(w http.ResponseWriter, req *http.Request) {
	users := []model.User{}

	users = decodeJSON(w, req, users)

	err = model.InsertUsers(users)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	response, err := json.Marshal(ServerCreatedMessage)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	w.Write(response)
}

func DeleteUsers(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		usersIds := []int{}

		err := json.NewDecoder(req.Body).Decode(&usersIds)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		err = model.DeleteUsers(usersIds)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		data, _ := json.Marshal(ServerDeletedMessage)

		w.Write(data)

	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func UpdateUser(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		ub := model.UpdateUserBody{}

		err := json.NewDecoder(req.Body).Decode(&ub)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		err = model.UpdateUser(ub)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			return
		}

		data, _ := json.Marshal(ServerUpdatedMessage)

		w.Write(data)

	}
}
