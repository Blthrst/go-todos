package handlers

import (
	"github.com/Blthrst/go-todos/model"
	"net/http"
	"encoding/json"
	"log"
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

func decodeJSON[W http.ResponseWriter, R *http.Request, T any](w W, req R, target T) T {
	err := json.NewDecoder(req.Body).Decode(&target)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		return 
	}

	return target
}