package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/r-malon/sgaf/db"
)

func createValor(w http.ResponseWriter, r *http.Request) error {
	valor := first(strconv.ParseInt(r.FormValue("valor"), 10, 64))
	dataInicio := ISO8601DateRegex.FindString(r.FormValue("data_inicio"))
	dataFim := ISO8601DateRegex.FindString(r.FormValue("data_fim"))
	itemID := first(strconv.ParseInt(r.FormValue("item_id"), 10, 64))

	if errors := validateValor(valor, dataInicio, dataFim, itemID); len(errors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(w).Encode(errors)
	}

	_, err := q.GetItem(ctx, itemID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return &ValidationError{
			Field:   "item_id",
			Message: "ITEM NÃO ENCONTRADO",
		}
	}

	valores, err := q.ListValorsByItem(ctx, itemID)
	if err == nil {
		for _, v := range valores {
			// Valor ativo encontrado
			if v.DataFim.String == "" {
				w.WriteHeader(http.StatusConflict)
				return &ValidationError{
					Field:   "item_id",
					Message: "JÁ EXISTE UM VALOR ATIVO PARA ESTE ITEM",
				}
			}
		}
	}

	return q.CreateValor(ctx, db.CreateValorParams{
		Valor:      valor,
		DataInicio: dataInicio,
		DataFim: sql.NullString{
			String: dataFim,
			Valid:  dataFim != "",
		},
		ItemID: itemID,
	})
}

func listValors(w http.ResponseWriter, r *http.Request) error {
	l, err := q.ListValors(ctx)
	if err != nil {
		return err
	}
	return tmpl.ExecuteTemplate(w, "listValors", l)
}

func updateValor(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return &ValidationError{Field: "id", Message: "ID INVÁLIDO"}
	}

	valor, err := q.GetValor(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return &ValidationError{Field: "id", Message: "VALOR NÃO ENCONTRADO"}
	}

	novoValor := first(strconv.ParseInt(r.FormValue("valor"), 10, 64))
	dataInicio := ISO8601DateRegex.FindString(r.FormValue("data_inicio"))
	dataFim := ISO8601DateRegex.FindString(r.FormValue("data_fim"))
	itemID := first(strconv.ParseInt(r.FormValue("item_id"), 10, 64))

	if errors := validateValor(novoValor, dataInicio, dataFim, itemID); len(errors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		return json.NewEncoder(w).Encode(errors)
	}

	// Verify if Item exists (if changed)
	if itemID != valor.ItemID {
		_, err := q.GetItem(ctx, itemID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return &ValidationError{
				Field:   "item_id",
				Message: "ITEM NÃO ENCONTRADO",
			}
		}

		// Verificar se já existe um valor ativo para o novo item
		valores, err := q.ListValorsByItem(ctx, itemID)
		if err == nil {
			for _, v := range valores {
				if v.ID != id && v.DataFim.String == "" { // Valor ativo encontrado
					w.WriteHeader(http.StatusConflict)
					return &ValidationError{
						Field:   "item_id",
						Message: "JÁ EXISTE UM VALOR ATIVO PARA ESTE ITEM",
					}
				}
			}
		}
	}

	return q.UpdateValor(ctx, db.UpdateValorParams{
		ID:         id,
		Valor:      novoValor,
		DataInicio: dataInicio,
		DataFim: sql.NullString{
			String: dataFim,
			Valid:  dataFim != "",
		},
		ItemID: itemID,
	})
}

func deleteValor(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return &ValidationError{Field: "id", Message: "ID INVÁLIDO"}
	}

	_, err = q.GetValor(ctx, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return &ValidationError{Field: "id", Message: "VALOR NÃO ENCONTRADO"}
	}

	return q.DeleteValor(ctx, id)
}
