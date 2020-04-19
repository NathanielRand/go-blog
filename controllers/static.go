package controllers

import "github.com/NathanielRand/go-blog/views"

type Static struct {
	Home	 *views.View
	Contact  *views.View
	Faq 	 *views.View
}

func NewStatic() *Static {
	return &Static{
		Home: views.NewView("materialize", "static/home"),
		Contact: views.NewView("materialize", "static/contact"),
		Faq: views.NewView("materialize", "static/faq"),
	}
}