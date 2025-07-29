package main

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/", templ.Handler(index()))
	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Println(err)
			return
		}
		text := r.FormValue("word-input")
		content := outputText(text)
		content.Render(r.Context(), w)
	})

	fmt.Println("Listening on :5000")
	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println("error port :5000")
	}
}
