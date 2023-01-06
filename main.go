package main

import (
	"classic"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type HangmanWeb struct {
	classic classic.HangManData
	Nletter string
}

func HandlePage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/home.html"))
	data := HangmanWeb{
		classic: classic.HangManData{
			Try:             "",
			Letter:          "",
			Randomword:      strings.ToUpper(classic.Randomword()),
			TotalTries:      10,
			NFormula:        0,
			Slice:           []string{},
			SliceRandomword: []string{},
			Boolean:         false,
		},
	}
	data.classic.NFormula = len(classic.Randomword())/2 - 1
	data.classic.Slice = make([]string, len(data.classic.Randomword))
	data.classic.SliceRandomword = make([]string, len(data.classic.Randomword))
	classic.PrintLettersInTheFullSlice(&data.classic)
	classic.Start(&data.classic)
	data.Nletter = classic.PrintNLetters(data.classic)
	t.Execute(w, data)
}
func main() {
	print("http://localhost:8080")
	http.HandleFunc("/", HandlePage)
	http.ListenAndServe(":8080", nil)
	fmt.Printf("Starting server at port 8080\n")
}
