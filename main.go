package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Client represents a customer
type Result struct {
	Name string
	Age  int
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	p := Result{Name: "Hasan Ayyad", Age: 27}
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, p)
}

func resultsHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	t, _ := template.ParseFiles("results.html")
	t.Execute(w, "")

	fmt.Fprintln(w, r.Form)
}

func main() {
	// Handling static assets
	files := http.FileServer(http.Dir("public"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// Routing requests
	http.HandleFunc("/form/", formHandler)
	http.HandleFunc("/results/", resultsHandler)

	// Starting server
	http.ListenAndServe(":8080", nil)
}
