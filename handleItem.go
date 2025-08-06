package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/r-malon/sgaf/db"
)

func createItem(w http.ResponseWriter, r *http.Request) error {
	descricao := r.FormValue("descricao")
	bandaMaxima := first(strconv.ParseInt(r.FormValue("banda_maxima"), 10, 64))
	bandaInstalada := first(strconv.ParseInt(r.FormValue("banda_instalada"), 10, 64))
	dataInstalacao := ISO8601DateRegex.FindString(r.FormValue("data_instalacao"))
	afID := first(strconv.ParseInt(r.FormValue("af_id"), 10, 64))
	localID := first(strconv.ParseInt(r.FormValue("local_id"), 10, 64))
	quantidade := first(strconv.ParseInt(r.FormValue("quantidade"), 10, 64))
	status := first(strconv.ParseBool(r.FormValue("status")))

	if errors := validateItem(descricao, bandaMaxima, bandaInstalada, dataInstalacao, quantidade, status); len(errors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(w).Encode(errors)
	}

	_, err := q.GetAF(ctx, afID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return &ValidationError{
			Field:   "af_id",
			Message: "AF NÃO ENCONTRADA",
		}
	}

	_, err = q.GetLocal(ctx, localID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return &ValidationError{
			Field:   "local_id",
			Message: "LOCAL NÃO ENCONTRADO",
		}
	}

	return q.CreateItem(ctx, db.CreateItemParams{
		Descricao:      descricao,
		BandaMaxima:    bandaMaxima,
		BandaInstalada: bandaInstalada,
		DataInstalacao: dataInstalacao,
		LocalID:        localID,
		AfID:           afID,
		Quantidade:     quantidade,
		Status:         status,
	})
}

func listItems(w http.ResponseWriter, r *http.Request) error {
	l, err := q.ListItems(ctx)
	if err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(w, "listItems", l)
}

func updateItem(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return &ValidationError{Field: "id", Message: "ID INVÁLIDO"}
	}

	item, err := q.GetItem(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return &ValidationError{Field: "id", Message: "ITEM NÃO ENCONTRADO"}
	}

	descricao := r.FormValue("descricao")
	bandaMaxima := first(strconv.ParseInt(r.FormValue("banda_maxima"), 10, 64))
	bandaInstalada := first(strconv.ParseInt(r.FormValue("banda_instalada"), 10, 64))
	dataInstalacao := ISO8601DateRegex.FindString(r.FormValue("data_instalacao"))
	localID := first(strconv.ParseInt(r.FormValue("local_id"), 10, 64))
	quantidade := first(strconv.ParseInt(r.FormValue("quantidade"), 10, 64))
	status := first(strconv.ParseBool(r.FormValue("status")))

	if errors := validateItem(descricao, bandaMaxima, bandaInstalada, dataInstalacao, quantidade, status); len(errors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(w).Encode(errors)
	}

	// Verify if Local exists (if changed)
	if localID != item.LocalID {
		_, err := q.GetLocal(ctx, localID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return &ValidationError{
				Field:   "local_id",
				Message: "LOCAL NÃO ENCONTRADO",
			}
		}
	}

	return q.UpdateItem(ctx, db.UpdateItemParams{
		ID:             id,
		Descricao:      descricao,
		BandaMaxima:    bandaMaxima,
		BandaInstalada: bandaInstalada,
		DataInstalacao: dataInstalacao,
		LocalID:        localID,
		Quantidade:     quantidade,
		Status:         status,
	})
}

func deleteItem(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		return err
	}
	return q.DeleteItem(ctx, id)
}
