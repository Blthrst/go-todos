package main

type Todo struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"isDone"`
	UserId      uint   `json:"userId"`
}

type User struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}
