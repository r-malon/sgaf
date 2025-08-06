package main

import (
	"context"
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"

	_ "github.com/joho/godotenv/autoload"
	"github.com/r-malon/sgaf/db"
	_ "modernc.org/sqlite"
)

var (
	tmpl             *template.Template
	ctx              context.Context
	dbconn           *sql.DB
	q                *db.Queries
	ISO8601DateRegex *regexp.Regexp
)

type errHandler func(http.ResponseWriter, *http.Request) error

func main() {
	defer dbconn.Close()

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/{$}", home)

	http.Handle("GET /local/{$}", errHandler(listLocals))
	http.Handle("POST /local/{$}", errHandler(createLocal))
	http.Handle("PUT /local/{id}", errHandler(updateLocal))
	http.Handle("DELETE /local/{id}", errHandler(deleteLocal))

	http.Handle("GET /af/{$}", errHandler(listAFs))
	http.Handle("POST /af/{$}", errHandler(createAF))
	http.Handle("PUT /af/{id}", errHandler(updateAF))
	http.Handle("DELETE /af/{id}", errHandler(deleteAF))

	http.Handle("GET /item/{$}", errHandler(listItems))
	http.Handle("POST /item/{$}", errHandler(createItem))
	http.Handle("PUT /item/{id}", errHandler(updateItem))
	http.Handle("DELETE /item/{id}", errHandler(deleteItem))

	http.Handle("GET /valor/{$}", errHandler(listValors))
	http.Handle("POST /valor/{$}", errHandler(createValor))
	http.Handle("PUT /valor/{id}", errHandler(updateValor))
	http.Handle("DELETE /valor/{id}", errHandler(deleteValor))

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

func first[T, U any](v T, d U) T {
	switch any(d).(type) {
	case error:
		log.Printf("%v	%v\n", v, d)
	}
	return v
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
	ISO8601DateRegex = regexp.MustCompile(`\d{4}-([0][1-9]|1[0-2])-([0][1-9]|[1-2]\d|3[01])`)
}
