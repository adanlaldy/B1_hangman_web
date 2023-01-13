package main

import (
	"classic"
	"html/template"
	"net/http"
	"strings"
)

type HangmanWeb struct {
	classic classic.HangManData
	//Hangman string
	//Tries   int
}

var Data = HangmanWeb{
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

func HandlePage(w http.ResponseWriter, r *http.Request) {
	if classic.IfZeroTry(&Data.classic) == true {
		return
	}
	Data.classic.Try = r.FormValue("input")
	Data.classic.Try = strings.ToUpper(Data.classic.Try)
	if classic.IfInputIsTheFullWord(&Data.classic) == true {
		return
	}
	classic.IfInputIsTrue(&Data.classic)
	if classic.IfSliceIsFull(&Data.classic) == true {
		return
	}
	t := template.Must(template.ParseFiles("./templates/home.html"))
	t.Execute(w, struct {
		Tries int
		Slice string
	}{
		Data.classic.TotalTries,
		classic.PrintSlice(&Data.classic),
	})
}

func main() {
	Data.classic.NFormula = len(classic.Randomword())/2 - 1
	Data.classic.Slice = make([]string, len(Data.classic.Randomword))
	Data.classic.SliceRandomword = make([]string, len(Data.classic.Randomword))
	classic.PrintLettersInTheFullSlice(&Data.classic)
	classic.Start(&Data.classic)
	classic.PrintNLetters(&Data.classic)
	print("http://localhost:8080")
	http.HandleFunc("/", HandlePage)
	http.ListenAndServe(":8080", nil)
}
