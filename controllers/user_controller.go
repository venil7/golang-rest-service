package controllers

import (
	"encoding/json"
	"fmt"
	"mywebproj/models"
	service "mywebproj/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func InsertUser(w http.ResponseWriter, r *http.Request) {
	conn, err := service.CreateConnection()
	switch {
	case err != nil:
		fmt.Fprintf(w, err.Error())
	case conn != nil:
		defer conn.Close()
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		user.ID, err = service.InsertUser(user)
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		json.NewEncoder(w).Encode(user)
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	conn, err := service.CreateConnection()
	switch {
	case err != nil:
		fmt.Fprintf(w, err.Error())
	case conn != nil:
		defer conn.Close()
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		user, err := service.GetUser(int64(id))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		}
		json.NewEncoder(w).Encode(user)
	}
}
