package controllers

import (
	"github.com/tingwei628/pgo/urlshortner/internal/project"
	"net/http"
	"text/template"
)

func GetIndex(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles(project.WorkingDir + "/internal/views/index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
