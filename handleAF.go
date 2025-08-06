package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/r-malon/sgaf/db"
)

func createAF(w http.ResponseWriter, r *http.Request) error {
	numero, _ := strconv.ParseInt(r.FormValue("numero"), 10, 64)
	fornecedor := r.FormValue("fornecedor")
	descricao := r.FormValue("descricao")
	dataInicio := ISO8601DateRegex.FindString(r.FormValue("data_inicio"))
	dataFim := ISO8601DateRegex.FindString(r.FormValue("data_fim"))
	status, _ := strconv.ParseBool(r.FormValue("status"))

	if errors := validateAF(numero, fornecedor, descricao, dataInicio, dataFim, status); len(errors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(w).Encode(errors)
	}

	return q.CreateAF(ctx, db.CreateAFParams{
		Numero:     numero,
		Fornecedor: fornecedor,
		Descricao:  descricao,
		DataInicio: dataInicio,
		DataFim:    dataFim,
		Status:     status,
	})
}

func listAFs(w http.ResponseWriter, r *http.Request) error {
	l, err := q.ListAFs(ctx)
	if err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(w, "listAFs", l)
}

func updateAF(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return &ValidationError{Field: "id", Message: "ID INVÁLIDO"}
	}

	af, err := q.GetAF(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return &ValidationError{Field: "id", Message: "AF NÃO ENCONTRADA"}
	}

	numero, _ := strconv.ParseInt(r.FormValue("numero"), 10, 64)
	fornecedor := r.FormValue("fornecedor")
	descricao := r.FormValue("descricao")
	dataInicio := ISO8601DateRegex.FindString(r.FormValue("data_inicio"))
	dataFim := ISO8601DateRegex.FindString(r.FormValue("data_fim"))
	status, _ := strconv.ParseBool(r.FormValue("status"))

	if errors := validateAF(numero, fornecedor, descricao, dataInicio, dataFim, status); len(errors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(w).Encode(errors)
	}

	return q.UpdateAF(ctx, db.UpdateAFParams{
		ID:         af.ID,
		Numero:     numero,
		Fornecedor: fornecedor,
		Descricao:  descricao,
		DataInicio: dataInicio,
		DataFim:    dataFim,
		Status:     status,
	})
}

func deleteAF(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		return err
	}
	return q.DeleteAF(ctx, id)
}
