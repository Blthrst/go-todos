package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"github.com/Blthrst/go-todos/model"
)

func UsersHandler(w http.ResponseWriter, req *http.Request) {

	switch req.Method {
	case "GET":
		{
			userId, err := url.ParseQuery(req.URL.RawQuery)

			if err != nil {
				panic(err.Error())
			}

			id, err := strconv.Atoi(userId["userId"][0])

			if err != nil {
				log.Fatal(err.Error())
			}

			u, err := model.GetUser(id)

			if err != nil {
				log.Fatal(err.Error())
			}

			data, err := json.Marshal(u)

			w.Write(data)

			if err != nil {
				log.Fatal(err.Error())
			}
		}

	case "POST":
		{

			u := model.User{}

			err := json.NewDecoder(req.Body).Decode(&u)

			if err != nil {
				log.Fatal(err.Error())
			}

			_, err = model.InsertUser(u)

			if err != nil {
				log.Fatal(err.Error())
			}

			w.WriteHeader(http.StatusCreated)
		}

	default: w.WriteHeader(http.StatusNotFound)
	}
}
