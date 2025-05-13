package main

import (
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type User struct {
	Name  string
	Email string
}

func home(w http.ResponseWriter, r *http.Request) {
	u := User{"John Doe", "john@gmail.com"}

	templates.ExecuteTemplate(w, "home.html", u) // the third parameter is the data to be passed to the template
}

func main() {
	templates = template.Must(template.ParseGlob("*.html")) // parses all html files in the current directory

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	})

	http.HandleFunc("/home", home)

	log.Fatal(http.ListenAndServe(":8000", nil)) // creates a server
}
