package main

import (
	"classic"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type HangmanWeb struct {
	classic classic.HangManData
}

var Data = HangmanWeb{
	classic: classic.HangManData{
		Try:             "",
		Letter:          "",
		Randomword:      "",
		Difficulty:      "",
		Difficultchoice: "",
		Name:            "",
		Points:          100,
		TotalTries:      10,
		NFormula:        0,
		Jose:            0,
		Slice:           []string{},
		SliceRandomword: []string{},
		SliceTries:      []string{},
		Scoreboard:      []string{},
		InputTrue:       false,
		InputTrue2:      false,
	},
}

func LevelPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/level.html"))
	Data.classic.Scoreboard = append(Data.classic.Scoreboard, "")
	if !(r.FormValue("name") == "" && r.FormValue("difficulty") == "") {
		Data.classic.Name = r.FormValue("name")
		Data.classic.Scoreboard[0] = Data.classic.Name
		Data.classic.Difficulty = r.FormValue("difficulty")
		Data.classic.Randomword = strings.ToUpper(classic.Randomword(&Data.classic))
		Data.classic.NFormula = len(classic.Randomword(&Data.classic))/2 - 1
		Data.classic.Slice = make([]string, len(Data.classic.Randomword))
		Data.classic.SliceRandomword = make([]string, len(Data.classic.Randomword))
		Data.classic.Scoreboard[2] = strconv.Itoa(100)
		Data.classic.Jose = 0
		classic.PrintLettersInTheFullSlice(&Data.classic)
		classic.Start(&Data.classic)
		classic.PrintNLetters(&Data.classic)
		http.Redirect(w, r, "/game", 303)
	}
	t.Execute(w, r)
}
func Restart() {
	Data.classic.Randomword = strings.ToUpper(classic.Randomword(&Data.classic))
	Data.classic.NFormula = len(classic.Randomword(&Data.classic))/2 - 1
	Data.classic.Slice = make([]string, len(Data.classic.Randomword))
	Data.classic.SliceRandomword = make([]string, len(Data.classic.Randomword))
	Data.classic.Jose = 0
	Data.classic.TotalTries = 10
	Data.classic.SliceTries = []string{}
	classic.PrintLettersInTheFullSlice(&Data.classic)
	classic.Start(&Data.classic)
	classic.PrintNLetters(&Data.classic)
}
func WinPage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/YouWin.html"))
	if r.FormValue("restart") != "" {
		http.Redirect(w, r, "/level", 303)
	}
	t.Execute(w, struct {
		Randomword string
		Scoreboard string
	}{
		Data.classic.Randomword,
		classic.PrintScoreboard(&Data.classic),
	})
	Restart()
}
func LoosePage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/GameOver.html"))
	if r.FormValue("restart") != "" {
		http.Redirect(w, r, "/level", 303)
	}
	t.Execute(w, struct {
		Randomword string
		Scoreboard string
	}{
		Data.classic.Randomword,
		classic.PrintScoreboard(&Data.classic),
	})
	Restart()
}
func GamePage(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/home.html"))
	if r.FormValue("input") != "" {
		Data.classic.InputTrue = false
		Data.classic.InputTrue2 = false
		if classic.IfZeroTry(&Data.classic) == true {
			http.Redirect(w, r, "/loose", 303)
		}
		Data.classic.Try = strings.ToUpper(r.FormValue("input"))
		if classic.IfInputIsTheFullWord(&Data.classic) == true {
			http.Redirect(w, r, "/win", 303)
		}
		if classic.Ifinputisthesame(&Data.classic) == true {
		} else {
			if classic.IfInputIsTrue(&Data.classic) == false {
				Data.classic.Jose++
				Data.classic.TotalTries--
				Data.classic.Points -= 10
				Data.classic.Scoreboard[2] = strconv.Itoa(Data.classic.Points)
				//Data.classic.Scoreboard2 = append(Data.classic.Scoreboard2, Data.classic.Scoreboard[0], Data.classic.Scoreboard[1], Data.classic.Scoreboard[2]+"\n")
				Data.classic.SliceTries = append(Data.classic.SliceTries, Data.classic.Try)
			} else if Data.classic.InputTrue == true {
				Data.classic.SliceTries = append(Data.classic.SliceTries, Data.classic.Try)
			}
		}
		if classic.IfSliceIsFull(&Data.classic) == true {
			http.Redirect(w, r, "/win", 303)
		}
	}
	t.Execute(w, struct {
		Tries      int
		Slice      string
		SliceTries []string
		Jose       int
		Scoreboard string
	}{
		Data.classic.TotalTries,
		classic.PrintSlice(&Data.classic),
		Data.classic.SliceTries,
		Data.classic.Jose,
		classic.PrintScoreboard(&Data.classic),
	})
}
func main() {
	fsJose := http.FileServer(http.Dir("jose"))
	fsCss := http.FileServer(http.Dir("static/css/"))
	fsSource := http.FileServer(http.Dir("static/source/"))
	fmt.Println("http://localhost:8080/level")
	http.HandleFunc("/level", LevelPage)
	http.HandleFunc("/game", GamePage)
	http.HandleFunc("/win", WinPage)
	http.HandleFunc("/loose", LoosePage)
	http.Handle("/jose/", http.StripPrefix("/jose", fsJose))
	http.Handle("/css/", http.StripPrefix("/css/", fsCss))
	http.Handle("/source/", http.StripPrefix("/source/", fsSource))
	http.ListenAndServe(":8080", nil)
}
