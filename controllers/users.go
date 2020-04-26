package controllers

import (
	"fmt"
	"net/http"

	"github.com/NathanielRand/go-blog/models"
	"github.com/NathanielRand/go-blog/views"
)

type Users struct {
	NewView *views.View
	us 		*models.UserService
}

type SignupForm struct {
	Username string `schema:"username"`
	Email    string `schema:"email"`
	Password string `scheam:"password"`
}

func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: 	views.NewView("materialize", "users/new"),
		us: 		us,
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

	user := models.User{
		Username: form.Username,
		Email:	  form.Email,
	}

	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "User is", user)
	fmt.Fprintln(w, "Username is", form.Username)
	fmt.Fprintln(w, "Email is ", form.Email)
	fmt.Fprintln(w, "Password is ", form.Password)
}
