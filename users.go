package main

import (
	"errors"
	//"fmt"
	"net/http"
	"net/url"
)

func GetUser(w http.ResponseWriter, req *http.Request) {

    _, err := url.ParseQuery(req.URL.RawQuery)

	if err != nil {
		errors.New("error when parse url query")
	}

	//get user id and find in database. import function from db

}