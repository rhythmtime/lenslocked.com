package controllers

import (
	"fmt"
	"lenslocked.com/views"
	"net/http"
	// "reflect"
	"github.com/gorilla/schema"
)

//NewUsers is used to create a new Users controller.
//This function will panic if the templates are not
//parsed correctly, and should only be used during
//initial setup.
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

// New is used to render the form where a user can
// create a new user account.
//
// GET /signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// struct tags point to fields in a struct
// these are used in json
// these require the reflect package
// the reflect package is for very advanced users!
// this is explained in gorilla schema docs
// the words following schema: refer to the name <input name= in the gohtml
type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Create is used to process the signup form when a user
// submits it. This is used to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	dec := schema.NewDecoder()
	var form SignupForm
	if err := dec.Decode(&form, r.PostForm); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}

// in new.gohtml Postform etc is looking at input name=""
// this is how to get data from a form in golang HTTP
// instead we are using gorilla schema to do this
//fmt.Fprintln(w, r.PostForm["email"])
//fmt.Fprintln(w, r.PostFormValue("email"))
