package main

import (
	"net/http"
	"strconv"

	"github.com/r-malon/sgaf/db"
)

func createItem(w http.ResponseWriter, r *http.Request) error {
	return q.CreateItem(ctx, db.CreateItemParams{
		r.FormValue("descricao"),
		first(strconv.ParseInt(r.FormValue("banda_maxima"), 10, 64)),
		first(strconv.ParseInt(r.FormValue("banda_instalada"), 10, 64)),
		ISO8601DateRegex.FindString(r.FormValue("data_instalacao")),
	})
}

func listItems(w http.ResponseWriter, r *http.Request) error {
	l, err := q.ListItems(ctx)
	tmpl.ExecuteTemplate(w, "listItems", l)
	return err
}

func updateItem(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)
	return q.UpdateItem(ctx, db.UpdateItemParams{
		r.FormValue("descricao"),
		first(strconv.ParseInt(r.FormValue("banda_maxima"), 10, 64)),
		first(strconv.ParseInt(r.FormValue("banda_instalada"), 10, 64)),
		ISO8601DateRegex.FindString(r.FormValue("data_instalacao")),
		id,
	})
}

func deleteItem(w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.ParseInt(r.PathValue("id"), 10, 64)
	return q.DeleteItem(ctx, id)
}
