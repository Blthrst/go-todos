package handlers

import (
	"github.com/Blthrst/go-todos/model"
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