package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Client represents a customer
type Client struct {
	Name string
	Age  int
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	p := Client{Name: "Hasan Ayyad", Age: 27}
	t, _ := template.ParseFiles("form.html")
	t.Execute(w, p)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Whoa, Go is neat!</h1>")
}

func main() {
	// Handling static assets
	files := http.FileServer(http.Dir("public"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// Routing requests
	http.HandleFunc("/", formHandler)
	http.HandleFunc("/result/", resultHandler)

	http.ListenAndServe(":8080", nil)
}
