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
	},
}

// A faire plus propre
// L'array qui stock les lettres déjà utilisés commence à beuguer quand il y a beaucoup de lettres ou quand on input des majuscules
func HandlePage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/home.html"))
	boolean := false
	for i := 0; i < len(Data.classic.SliceTries); i++ {
		if Data.classic.SliceTries[i] == r.FormValue("input") {
			boolean = true
		}
	}
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
			Data.classic.SliceTries = append((Data.classic.SliceTries), r.FormValue("input"))
		} else if classic.IfInputIsTrue(&Data.classic) == true && boolean == false {
			Data.classic.SliceTries = append((Data.classic.SliceTries), r.FormValue("input"))
		} else if classic.IfInputIsTrue(&Data.classic) == true || classic.IfInputIsTrue(&Data.classic) == false && boolean == false {
			fmt.Println("you cannot enter twice the same letter")
		}

		if classic.IfSliceIsFull(&Data.classic) == true {
			return
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
