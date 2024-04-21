package model

type Todo struct {
	Id          int   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
	UserId      int   `json:"userId"`
}

type User struct {
	Id   int   `json:"id"`
	Name string `json:"name"`
}
