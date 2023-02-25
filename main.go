package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var username string = "Naman Bhardwaj"
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hi! "+username+"... Welcome to <font color = 'orange'>Firenos</font></h1>")
}

func Contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Hi! To get in touch please send an email to <a href = \"mailto:support@lenslocked.com\">support@lenslocked.com</a>")
}

func Faq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Faq Page</h1><p>Under Development</p>")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>404 Page Not found!</h1>")
}

func main() {
	template.New("Blah")
	r := httprouter.New()
	r.GET("/", Home)
	r.GET("/contact", Contact)
	r.GET("/faq", Faq)
	r.NotFound = http.HandlerFunc(notFound)
	http.ListenAndServe(":3000", r)
}
