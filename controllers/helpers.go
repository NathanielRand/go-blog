package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

func parseForm(r *http.Request, dst interface{}) error {
	// Parse the HTML form that was submitted
	// and store the data in the PostForm field of the http.Request.
	if err := r.ParseForm(); err != nil {
		return err
	}

	// Initialize our decoder.
	dec := schema.NewDecoder()	

	// Call the Decode method on our decode and
	// pass in the SignupForm as our destination and
	// r.PostForm as our data source. Check for errors.
	if err := dec.Decode(dst, r.PostForm); err != nil {
		return err
	}

	// Return nil when no errors occur.
	return nil
}