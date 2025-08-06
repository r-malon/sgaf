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
	if err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(w, "listLocals", l)
}

func updateLocal(w http.ResponseWriter, r *http.Request) error {
	return q.UpdateLocal(ctx, db.UpdateLocalParams{
		r.FormValue("nome"),
		first(strconv.ParseInt(r.PathValue("id"), 10, 64)),
	})
}

func deleteLocal(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		return err
	}
	return q.DeleteLocal(ctx, id)
}
