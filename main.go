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
		SliceTries:      []string{},
		Boolean:         false,
		Boolean2:        false,
	},
}

// A faire plus propre
func HandlePage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/game.html"))
	if r.FormValue("input") != "" {
		Data.classic.Boolean = false
		Data.classic.Boolean2 = false
		if classic.IfZeroTry(&Data.classic) == true {
			return
		}
		Data.classic.Try = strings.ToUpper(r.FormValue("input"))
		if classic.IfInputIsTheFullWord(&Data.classic) == true {
			return
		}
		if classic.Ifinputisthesame(&Data.classic) == true {
		} else {
			if classic.IfInputIsTrue(&Data.classic) == false {
				Data.classic.TotalTries--
				Data.classic.SliceTries = append((Data.classic.SliceTries), Data.classic.Try)
			} else if classic.IfInputIsTrue(&Data.classic) == true {
				Data.classic.SliceTries = append((Data.classic.SliceTries), Data.classic.Try)
			}
			if classic.IfSliceIsFull(&Data.classic) == true {
				return
			}
		}
	}
	t.Execute(w, struct {
		Tries      int
		Slice      string
		SliceTries []string
	}{
		Data.classic.TotalTries,
		classic.PrintSlice(&Data.classic),
		Data.classic.SliceTries,
	})
}
func main() {
	fs := http.FileServer(http.Dir("templates/"))
	print("http://localhost:8080")
	Data.classic.NFormula = len(classic.Randomword())/2 - 1
	Data.classic.Slice = make([]string, len(Data.classic.Randomword))
	Data.classic.SliceRandomword = make([]string, len(Data.classic.Randomword))
	classic.PrintLettersInTheFullSlice(&Data.classic)
	classic.Start(&Data.classic)
	classic.PrintNLetters(&Data.classic)
	http.HandleFunc("/", HandlePage)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}
