package main

import (
	"github.com/gorilla/mux"
	"lenslocked.com/controllers"
	"net/http"
)

func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()

	r := mux.NewRouter()

	//below we are passing the function itself not the results of the function
	r.Handle("/", staticC.Home).Methods("Get")
	r.Handle("/contact", staticC.Contact).Methods("Get")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":3000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
