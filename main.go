package main

import (
	"fmt"	
	"net/http"
	
	"github.com/NathanielRand/go-blog/controllers"

	"github.com/gorilla/mux"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Not Found, whomp...</h1>")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/favicon.ico")
}

var nf http.Handler = http.HandlerFunc(notFound)

func main() {
	
	// Controllers
	usersC := controllers.NewUsers()
	staticC := controllers.NewStatic()

	// Gorilla Mux Router
	r := mux.NewRouter()

	// 404 Not Found
	r.NotFoundHandler = nf

	// Favicon
	r.HandleFunc("/favicon.ico", faviconHandler)

	// Assest Routes
	assetHandler := http.FileServer(http.Dir("./assets/"))
	assetHandler = http.StripPrefix("/assets/", assetHandler)
	r.PathPrefix("/assets/").Handler(assetHandler)
	
	// Routes - Static
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.Handle("/faq", staticC.Faq).Methods("GET")
	
	// Routes - Users
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")

	// Start web server, listening on port 3000.
	http.ListenAndServe(":3000", r)
}
