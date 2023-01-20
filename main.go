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
		Name:            "",
		Randomword:      strings.ToUpper(classic.Randomword()),
		TotalTries:      10,
		NFormula:        0,
		Jose:            0,
		Slice:           []string{},
		SliceRandomword: []string{},
		SliceTries:      []string{},
		Boolean:         false,
		Boolean2:        false,
	},
}

func HandlePage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/home.html"))
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
				Data.classic.Jose++
				Data.classic.TotalTries--
				Data.classic.SliceTries = append((Data.classic.SliceTries), Data.classic.Try)
			} else if Data.classic.Boolean == true {
				Data.classic.SliceTries = append((Data.classic.SliceTries), Data.classic.Try)
			}
		}
		if classic.IfSliceIsFull(&Data.classic) == true {
			return
		}
	}
	t.Execute(w, struct {
		Tries      int
		Slice      string
		SliceTries []string
		Jose       int
	}{
		Data.classic.TotalTries,
		classic.PrintSlice(&Data.classic),
		Data.classic.SliceTries,
		Data.classic.Jose,
	})
}

func main() {
	fs := http.FileServer(http.Dir("jose"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
	print("http://localhost:8080")
	Data.classic.NFormula = len(classic.Randomword())/2 - 1
	Data.classic.Slice = make([]string, len(Data.classic.Randomword))
	Data.classic.SliceRandomword = make([]string, len(Data.classic.Randomword))
	classic.PrintLettersInTheFullSlice(&Data.classic)
	classic.Start(&Data.classic)
	classic.PrintNLetters(&Data.classic)
	Data.classic.Jose = 0
	/*t := template.Must(template.ParseFiles("./templates/home.html"))
	Data.classic.Name = (r.FormValue("name"))
	t.Execute(w, struct {})*/
	http.HandleFunc("/", HandlePage)
	http.ListenAndServe(":8080", nil)
}
