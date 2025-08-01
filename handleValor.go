package main

import (
	"net/http"
	"strconv"
	"database/sql"

	"github.com/r-malon/sgaf/db"
)

func createValor(w http.ResponseWriter, r *http.Request) error {
	return q.CreateValor(ctx, db.CreateValorParams{
		first(strconv.ParseInt(r.FormValue("valor"), 10, 64)),
		ISO8601DateRegex.FindString(r.FormValue("data_inicio")),
		sql.NullString{
			r.FormValue("data_fim"),
			ISO8601DateRegex.MatchString(r.FormValue("data_fim")),
		},
	})
}

func listValors(w http.ResponseWriter, r *http.Request) error {
	l, err := q.ListValors(ctx)
	tmpl.ExecuteTemplate(w, "listValors", l)
	return err
}

func updateValor(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)
	return q.UpdateValor(ctx, db.UpdateValorParams{
		first(strconv.ParseInt(r.FormValue("valor"), 10, 64)),
		ISO8601DateRegex.FindString(r.FormValue("data_inicio")),
		sql.NullString{
			r.FormValue("data_fim"),
			ISO8601DateRegex.MatchString(r.FormValue("data_fim")),
		},
		id,
	})
}

func deleteValor(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)
	return q.DeleteValor(ctx, id)
}
