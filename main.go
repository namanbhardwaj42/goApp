package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome to <font color = 'orange'>Firenos</font></h1>")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "Hi! To get in touch please send an email to <a href = \"mailto:support@lenslocked.com\">support@lenslocked.com</a>")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Home)
	r.HandleFunc("/contact", Contact)
	//http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", r)
}
