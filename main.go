package main

import (
	"classic"
	"fmt"
	"html/template"
	"net/http"
)

func HandlePage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/home.html"))
	input := "t pa bo"
	data := classic.HangManData{
		Try:             "",
		Letter:          input,
		Randomword:      "",
		TotalTries:      0,
		NFormula:        0,
		Slice:           []string{},
		SliceRandomword: []string{},
		Boolean:         true,
	}
	t.Execute(w, data)
}

func main() {
	print("http://localhost:8080")
	http.HandleFunc("/", HandlePage)
	http.ListenAndServe(":8080", nil)
	fmt.Printf("Starting server at port 8080\n")
}
