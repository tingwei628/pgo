package controllers

import (
	"database/sql"
	"github.com/tingwei628/pgo/urlshortener/internal/project"
	"github.com/tingwei628/pgo/urlshortener/internal/repository"
	"github.com/tingwei628/pgo/urlshortener/internal/url"
	"net/http"
	"strings"
	"text/template"
	"time"
)

const (
	HTTP_PREFIX  = "http://"
	HTTPS_PREFIX = "https://"
	DEFAULT_PATH = "/urlshortener"
)

func GetShorten(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// get url from  form-data
		originalURL := r.FormValue("url")

		// validate url
		if strings.TrimSpace(originalURL) == "" {
			http.Error(w, "url required", http.StatusBadRequest)
			return
		}
		if strings.HasPrefix(originalURL, HTTP_PREFIX) == false &&
			strings.HasPrefix(originalURL, HTTPS_PREFIX) == false {
			originalURL = HTTPS_PREFIX + originalURL
		}

		shortURL := url.Shorten(originalURL)

		// time to utc+0
		if err := repository.StoreURL(db, shortURL, originalURL, time.Now().UTC()); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// ShortURL in shorten.html
		data := map[string]string{
			"ShortURL": shortURL,
		}

		tmpl, err := template.ParseFiles(project.WorkingDir + "/internal/views/shorten.html")

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// fill html with data
		if err = tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetProxy(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//shortUrl := r.URL.Path
		shortUrl := strings.TrimPrefix(r.URL.Path, DEFAULT_PATH+"/")
		if strings.TrimSpace(r.URL.Path) == "" || strings.TrimSpace(shortUrl) == "" {
			http.Error(w, "url not provided", http.StatusBadRequest)
			return
		}
		originalUrl, err := repository.GetOriginalURL(db, shortUrl)
		if err != nil || originalUrl == "" {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		// redirect
		http.Redirect(w, r, originalUrl, http.StatusPermanentRedirect)

	}
}
