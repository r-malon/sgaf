package main

import (
	"net/http"
	"html/template"
	"database/sql"
	"context"
	"fmt"
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "modernc.org/sqlite"
	"github.com/r-malon/sgaf/db"
)

var tmpl *template.Template
var ctx context.Context
var dbconn *sql.DB
var q *db.Queries

func main() {
	defer dbconn.Close()

	q = db.New(dbconn)
	q.CreateLocal(ctx, "CMEI")
	fmt.Print("rteete")
	l, err := q.ListLocals(ctx)
	fmt.Printf("%+v %v", l, err)
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	r.Get("/", home)
/*
	r.Get("/item", listItems)
	r.Get("/item/{id:\\d}", getItem)
	r.Post("/item", addItem)
	r.Delete("/item/{id:\\d}", delItem)
*/
	http.ListenAndServe(":3000", r)
}

func init() {
	var err error
	ctx = context.TODO()
	dbconn, err = sql.Open("sqlite", ":memory:")

	if err != nil {
		log.Fatal(err)
	}
	if err = dbconn.Ping(); err != nil {
		log.Fatal(err)
	}

	tmpl, _ = template.ParseGlob("templates/*.html")
}

func home(w http.ResponseWriter, r *http.Request) {
	l, _ := q.ListLocals(ctx)
	fmt.Fprintf(w, "%+v", l)
}

