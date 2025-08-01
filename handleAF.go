package main

import (
	"net/http"
	"strconv"

	"github.com/r-malon/sgaf/db"
)

func createAF(w http.ResponseWriter, r *http.Request) error {
	return q.CreateAF(ctx, db.CreateAFParams{
		first(strconv.ParseInt(r.FormValue("numero"), 10, 64)),
		r.FormValue("fornecedor"),
		r.FormValue("descricao"),
		ISO8601DateRegex.FindString(r.FormValue("data_inicio")),
		ISO8601DateRegex.FindString(r.FormValue("data_fim")),
		first(strconv.ParseBool(r.FormValue("status"))),
	})
}

func listAFs(w http.ResponseWriter, r *http.Request) error {
	l, err := q.ListAFs(ctx)
	tmpl.ExecuteTemplate(w, "listAFs", l)
	return err
}

func updateAF(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)
	return q.UpdateAF(ctx, db.UpdateAFParams{
		first(strconv.ParseInt(r.FormValue("numero"), 10, 64)),
		r.FormValue("fornecedor"),
		r.FormValue("descricao"),
		ISO8601DateRegex.FindString(r.FormValue("data_inicio")),
		ISO8601DateRegex.FindString(r.FormValue("data_fim")),
		first(strconv.ParseBool(r.FormValue("status"))),
		id,
	})
}

func deleteAF(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)
	return q.DeleteAF(ctx, id)
}
