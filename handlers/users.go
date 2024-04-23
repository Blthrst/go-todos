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

	err := json.NewDecoder(req.Body).Decode(&users)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	_, err = model.InsertUsers(users)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	response, err := json.Marshal(MessageFromServer{
		Message: "Succesfully created",
		Status:  http.StatusCreated,
	})

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return
	}

	w.Write(response)
}
