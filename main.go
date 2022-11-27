package main

import (
	"fmt"
	"net/http"
	template2 "text/template"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home", "")
}

func RenderTemplate(w http.ResponseWriter, template string, data interface{}) {
	t, err := template2.ParseFiles("./templates/" + template + ".html")
	if err != nil {
		fmt.Println("ERROR : Can't render template", err)
	}
	t.Execute(w, "")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil)
	print("http://localhost:8080")
}
