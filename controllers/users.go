package controllers

import (
	"fmt"
	"net/http"

	"github.com/NathanielRand/go-blog/views"
)

type Users struct {
	NewView *views.View
}

type SignupForm struct {
	Username string `schema:"username"`
	Email    string `schema:"email"`
	Password string `scheam:"password"`
}

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("materialize", "users/new"),
	}
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

// Tries to create a new user account.
//
// POST /signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm

	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	fmt.Fprintln(w, "Username is", form.Username)
	fmt.Fprintln(w, "Email is ", form.Email)
	fmt.Fprintln(w, "Password is ", form.Email)
}
