package controllers

import (
	"net/http"
	
	"github.com/NathanielRand/go-blog/views"
)

type Users struct {
	NewView *views.View
}

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("materialize", "views/users/new.gohtml"),
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