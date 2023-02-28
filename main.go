package main

import (
	"fmt"
	"main/views"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	//homeTemplate    *template.Template
	//contactTemplate *template.Template
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := homeView.Template.ExecuteTemplate(w, homeView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := contactView.Template.ExecuteTemplate(w, contactView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func Faq(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Faq Page</h1><p>Under Development</p>")
}

func main() {
	/*
		var err error
		homeTemplate, err = template.ParseFiles(
			"views/home.gohtml",
			"views/layouts/footer.gohtml",
		)
		if err != nil {
			panic(err)
		}
		contactTemplate, err = template.ParseFiles(
			"views/contact.gohtml",
			"views/layouts/footer.gohtml",
		)
		if err != nil {
			panic(err)
		}
	*/

	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	r.HandleFunc("/faq", Faq)
	http.ListenAndServe(":3000", r)
}
