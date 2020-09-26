package main

import (
	"log"
	controllers "mywebproj/controllers"
	http "net/http"

	// auth "github.com/goji/httpauth"
	mux "github.com/gorilla/mux"
	genv "github.com/sakirsensoy/genv"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
)

func main() {
	listen := genv.Key("LISTEN").Default(":8080").String()

	r := mux.NewRouter()
	// r.Use(auth.SimpleBasicAuth("user", "password"))
	r.HandleFunc("/user", controllers.InsertUser).Methods("POST", "OPTIONS")
	r.HandleFunc("/user/{id}", controllers.GetUser).Methods("GET", "OPTIONS")

	http.Handle("/", r)
	log.Printf("listening on %s", listen)
	http.ListenAndServe(listen, nil)
}
