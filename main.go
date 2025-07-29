package main

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"github.com/r-malon/sgaf/db"
	_ "modernc.org/sqlite"
)

var (
	tmpl   *template.Template
	ctx    context.Context
	dbconn *sql.DB
	q      *db.Queries
)

type errHandler func(http.ResponseWriter, *http.Request) error

func main() {
	defer dbconn.Close()

	l, err := q.ListLocals(ctx)
	log.Printf("%+v %v\n", l, err)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/{$}", home)

	http.Handle("GET /local/{$}", errHandler(listLocals))
	http.Handle("POST /local/{$}", errHandler(createLocal))
	http.Handle("PUT /local/{id}", errHandler(updateLocal))
	http.Handle("DELETE /local/{id}", errHandler(deleteLocal))

	log.Fatal(http.ListenAndServe(os.Getenv("ADDR"), nil))
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "home", nil)
}

func (fn errHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func init() {
	var err error
	dbconn, err = sql.Open("sqlite", os.Getenv("DB_PATH"))

	if err != nil {
		log.Fatal(err)
	}

	if err = dbconn.Ping(); err != nil {
		log.Fatal(err)
	}

	q = db.New(dbconn)
	ctx = context.TODO()
	tmpl = template.Must(template.ParseGlob("templates/*.html.tmpl"))
}
