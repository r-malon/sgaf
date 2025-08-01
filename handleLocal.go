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
	return q.UpdateLocal(ctx, db.UpdateLocalParams{
		r.FormValue("nome"),
		first(strconv.ParseInt(r.PathValue("id"), 10, 64)),
	})
}

func deleteLocal(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)
	return q.DeleteLocal(ctx, id)
}
