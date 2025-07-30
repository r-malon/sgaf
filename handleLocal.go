package main

import (
	"net/http"
	"strconv"

	"github.com/r-malon/sgaf/db"
)

func createLocal(w http.ResponseWriter, r *http.Request) error {
	return q.CreateLocal(ctx, r.FormValue("nome"))
}

func listLocals(w http.ResponseWriter, r *http.Request) error {
	l, err := q.ListLocals(ctx)
	tmpl.ExecuteTemplate(w, "listLocals", l)
	return err
}

func updateLocal(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(r.PathValue("id"))
	nome := r.FormValue("nome")
	data := db.UpdateLocalParams{nome, int64(id)}
	return q.UpdateLocal(ctx, data)
}

func deleteLocal(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(r.PathValue("id"))
	return q.DeleteLocal(ctx, int64(id))
}
