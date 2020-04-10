package main

import (
	"fmt"	
	"net/http"

	"github.com/NathanielRand/go-blog/views"
	"github.com/NathanielRand/go-blog/controllers"

	"github.com/gorilla/mux"
)

var (
	homeView *views.View
	contactView *views.View
	// signupView *views.View
)


func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>FAQ</h1>")
}

// func signup(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	must(signupView.Render(w, nil))
// }

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Not Found, whomp...</h1>")
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/favicon.ico")
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

var nf http.Handler = http.HandlerFunc(notFound)

func main() {
	// Parse home template file.
	homeView = views.NewView("materialize", "views/home.gohtml")
	contactView = views.NewView("materialize", "views/contact.gohtml")
	// signupView = views.NewView("materialize", "views/users/new.gohtml")

	usersC := controllers.NewUsers()

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

	// Routes
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", faq)
	r.HandleFunc("/signup", usersC.New)

	// Start web server, listening on port 3000.
	http.ListenAndServe(":3000", r)
}
