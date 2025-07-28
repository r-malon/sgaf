package main

import (
	"net/http"
	"html/template"
	"database/sql"
	"context"
	"strconv"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
	"github.com/r-malon/sgaf/db"
)

var (
	tmpl *template.Template
	ctx context.Context
	dbconn *sql.DB
	q *db.Queries
)

func main() {
	defer dbconn.Close()

	q = db.New(dbconn)
	q.CreateLocal(ctx, "CMEI")
	l, err := q.ListLocals(ctx)
	fmt.Printf("%+v %v\n", l, err)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/{$}", home)

	http.HandleFunc("GET /local/", listLocals)
	http.HandleFunc("POST /local/", createLocal)
	http.HandleFunc("PUT /local/", updateLocal)
	http.HandleFunc("DELETE /local/", deleteLocal)

	http.ListenAndServe(":3000", nil)
}

func init() {
	var err error
	ctx = context.TODO()
	dbconn, err = sql.Open("sqlite", "test.db")

	if err != nil {
		log.Fatal(err)
	}
	if err = dbconn.Ping(); err != nil {
		log.Fatal(err)
	}

	tmpl, _ = template.ParseGlob("templates/*.html")
}

func home(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "home.html", nil)
}

func listLocals(w http.ResponseWriter, r *http.Request) {
	l, _ := q.ListLocals(ctx)
	tmpl.ExecuteTemplate(w, "listLocals", l)
}

func deleteLocal(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	if err := q.DeleteLocal(ctx, int64(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func createLocal(w http.ResponseWriter, r *http.Request) {
	if err := q.CreateLocal(ctx, r.FormValue("nome")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func updateLocal(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("id"))
	nome := r.FormValue("nome")
	data := db.UpdateLocalParams{nome, int64(id)}
	if err := q.UpdateLocal(ctx, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
