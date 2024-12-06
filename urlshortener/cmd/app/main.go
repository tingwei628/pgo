package main

import (
	"database/sql"
	"github.com/tingwei628/pgo/urlshortner/internal/controllers"
	"github.com/tingwei628/pgo/urlshortner/internal/middleware"
	"github.com/tingwei628/pgo/urlshortner/internal/project"
	"github.com/tingwei628/pgo/urlshortner/internal/repository"
	"log"
	_ "modernc.org/sqlite"
	"net/http"
	"time"
)

func main() {

	fs := http.FileServer(http.Dir("views"))
	http.Handle("/views/", http.StripPrefix("/views/", fs))
	db, err := sql.Open("sqlite", project.WorkingDir+"/internal/db.sqlite")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := repository.CreateTable(db); err != nil {
		log.Fatal(err)
	}
	// ttl = 5mins
	go repository.SetTTL(db, time.Minute*5)

	mux := http.NewServeMux()

	mux.HandleFunc(controllers.DEFAULT_PATH+"/{tail}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetProxy(db)(w, r)
	})

	mux.HandleFunc(controllers.DEFAULT_PATH, func(w http.ResponseWriter, r *http.Request) {
		controllers.GetIndex(w, r)
	})

	mux.HandleFunc(controllers.DEFAULT_PATH+"/shorten", controllers.GetShorten(db))

	wrappedMux := middleware.RateLimit(mux)

	log.Fatal(http.ListenAndServe(":8002", wrappedMux))
}
