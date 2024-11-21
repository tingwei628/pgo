package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tingwei628/pgo/urlshortner/internal/controllers"
	"github.com/tingwei628/pgo/urlshortner/internal/middleware"
	"github.com/tingwei628/pgo/urlshortner/internal/repository"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	workingDir, _ := os.Getwd() // get root directory

	fs := http.FileServer(http.Dir("views"))
	http.Handle("/views/", http.StripPrefix("/views/", fs))

	db, err := sql.Open("sqlite3", workingDir+"/urlshortner/internal/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := repository.CreateTable(db); err != nil {
		log.Fatal(err)
	}
	// ttl = 20s
	go repository.SetTTL(db, time.Second*20)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			controllers.GetIndex(w, r)
		} else {
			// shorten url
			controllers.GetProxy(db)(w, r)
		}
	})
	mux.HandleFunc("/shorten", controllers.GetShorten(db))

	wrappedMux := middleware.RateLimit(mux)

	log.Fatal(http.ListenAndServe(":8080", wrappedMux))
}
