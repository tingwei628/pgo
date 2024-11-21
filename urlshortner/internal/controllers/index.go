package controllers

import (
	"net/http"
	"os"
	"text/template"
)

func GetIndex(w http.ResponseWriter, _ *http.Request) {
	workingDir, _ := os.Getwd() // get root directory
	tmpl, err := template.ParseFiles(workingDir + "/urlshortner/internal/views/index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
