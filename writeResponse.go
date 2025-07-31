//go:build exclude
package main

import (
	"net/http"
	"encoding/json"
	"encoding/csv"
)

func writeResponse(w http.ResponseWriter, contentType string, data any) {
	switch r.Header.Get("Content-Type") {
	default:
		w.Header().Set("Content-Type", "text/html")
		tmpl.ExecuteTemplate(w, templateName, data)
	case "application/json":
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	case "text/plain":
		w.Header().Set("Content-Type", "text/plain")
		n := csv.NewWriter(w)
		n.Comma = '\t'
		n.Write(data)
	}
}
