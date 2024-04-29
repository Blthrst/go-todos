package model

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
	UserId      int    `json:"userId"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type MessageFromServer struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type UpdateUserBody struct {
	Id   int  `json:"id"`
	User User `json:"user"`
}

type CompleteTodoBody struct {
	Id int `json:"id"`
}