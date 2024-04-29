package handlers

import (
	"encoding/json"
	"github.com/Blthrst/go-todos/model"
	"log"
	"net/http"
)

var ServerCreatedMessage = model.MessageFromServer{
	Message: "successfully created",
	Status:  http.StatusCreated,
}

var ServerUpdatedMessage = model.MessageFromServer{
	Message: "successfully updated",
	Status:  http.StatusOK,
}

var ServerDeletedMessage = model.MessageFromServer{
	Message: "successfully deleted",
	Status:  http.StatusOK,
}

func sendServerMessage(w http.ResponseWriter, msg model.MessageFromServer) {
	data, _ := json.Marshal(msg)

	w.Write(data)
}

func decodeJSON[T any](w http.ResponseWriter, req *http.Request, target T) T {
	err := json.NewDecoder(req.Body).Decode(&target)

	if err != nil {
		writeErrorAndLog(w, err)
	}

	return target
}

func writeErrorAndLog(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusBadRequest)
	log.Println(err.Error())
}
