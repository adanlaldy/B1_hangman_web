package main

import (
	"classic"
	"html/template"
	"net/http"
	"strings"
)

type HangmanWeb struct {
	classic classic.HangManData
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
	t := template.Must(template.ParseFiles("./templates/home.html"))
	if r.FormValue("input") != "" {
		Data.classic.Boolean = false
		if classic.IfZeroTry(&Data.classic) == true {
			return
		}
		Data.classic.Try = strings.ToUpper(r.FormValue("input"))
		if classic.IfInputIsTheFullWord(&Data.classic) == true {
			return
		}
		if classic.IfInputIsTrue(&Data.classic) == false {
			Data.classic.TotalTries--
		}
		if classic.IfSliceIsFull(&Data.classic) == true {
			return
		}
	}
	t.Execute(w, struct {
		Tries int
		Slice string
	}{
		Data.classic.TotalTries,
		classic.PrintSlice(&Data.classic),
	})
}

func main() {
	print("http://localhost:8080")
	Data.classic.NFormula = len(classic.Randomword())/2 - 1
	Data.classic.Slice = make([]string, len(Data.classic.Randomword))
	Data.classic.SliceRandomword = make([]string, len(Data.classic.Randomword))
	classic.PrintLettersInTheFullSlice(&Data.classic)
	classic.Start(&Data.classic)
	classic.PrintNLetters(&Data.classic)
	http.HandleFunc("/", HandlePage)
	http.ListenAndServe(":8080", nil)
}
