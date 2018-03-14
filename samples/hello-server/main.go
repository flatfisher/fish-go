package main

import (
	"html/template"
	"net/http"
)

// Page is response.
type Page struct {
	Title string
	Count int
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"Hello World!!.", 1}
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe("localhost:8000", nil)
}
