package handlers

type MessageFromServer struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
