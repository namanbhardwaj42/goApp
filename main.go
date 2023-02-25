package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Welcome to <font color = 'red'>Firenos!</font></h1>")
	} else if r.URL.Path == "/contact" || r.URL.Path == "/contact/" {
		fmt.Fprint(w, "To get in touch please send an email to <a href = \"mailto:support@lenslocked.com\">support@lenslocked.com</a>")
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "<h1>We couldn't find the page you were looking for :( </h1><p>Please email us if you keep being sent to an invalid page</p>")
	}
}
func main() {
	mux := &http.ServeMux{}
	mux.HandleFunc("/", handlerFunc)
	//http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", mux)
}
