package main

import (
	//"fmt"
	"net/http"
)

func main() {

	InitSecrets()

	http.HandleFunc("/users/", GetUser)

    http.ListenAndServe("localhost:8090", nil)
}