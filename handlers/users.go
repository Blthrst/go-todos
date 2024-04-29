package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/Blthrst/go-todos/model"
)

func GetAllUsers(w http.ResponseWriter, req *http.Request) {
	users, err := model.GetAllUsers()

	if err != nil {
		writeErrorAndLog(w, err)
	}

	data, _ := json.Marshal(users)

	w.Write(data)
}

func GetOneUser(w http.ResponseWriter, req *http.Request) {

	userId, err := url.ParseQuery(req.URL.RawQuery)

	if err != nil {
		writeErrorAndLog(w, err)
	}

	id, err := strconv.Atoi(userId["userId"][0])

	if err != nil {
		writeErrorAndLog(w, err)
	}

	u, err := model.GetUser(id)

	if err != nil {
		writeErrorAndLog(w, err)
	}

	data, _ := json.Marshal(u)

	w.Write(data)

}

func CreateUsers(w http.ResponseWriter, req *http.Request) {
	users := []model.User{}

	users = decodeJSON(w, req, users)

	err := model.InsertUsers(users)

	if err != nil {
		writeErrorAndLog(w, err)
	}

	sendServerMessage(w, ServerCreatedMessage)
}

func DeleteUsers(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		usersIds := []int{}

		usersIds = decodeJSON(w, req, usersIds)

		err := model.DeleteUsers(usersIds)

		if err != nil {
			writeErrorAndLog(w, err)
		}

		sendServerMessage(w, ServerDeletedMessage)

	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func UpdateUser(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		ub := model.UpdateUserBody{}

		ub = decodeJSON(w, req, ub)

		err := model.UpdateUser(ub)

		if err != nil {
			writeErrorAndLog(w, err)
		}

		sendServerMessage(w, ServerUpdatedMessage)

	}
}