package main

import (
	"fmt"
	hangmanclassic "hangman"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "home", hangmanclassic.Structure)
}

func RenderTemplate(w http.ResponseWriter, template string, data interface{}) {
	t, err := template2.ParseFiles("./templates/" + template + ".html")
	if err != nil {
		fmt.Println("ERROR : Can't render template", err)
	}
	t.Execute(w, data)
}

func main() {
	hangmanclassic.Structure()
	http.HandleFunc("/", homeHandler)
	http.ListenAndServe(":8080", nil)
	fmt.Print("http://localhost:8080/%22")
}
