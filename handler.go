package main

import (
	"net/http"
)

func generateHandler(arcs map[string]Arc, templateFilename string) http.HandlerFunc {
	tmpl, err := generateTemplate(templateFilename)
	if err != nil {
		return nil
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/intro", http.StatusMovedPermanently)
	})

	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if arc, ok := arcs[path[1:]]; ok {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			tmpl.Execute(w, arc)
		} else {
			mux.ServeHTTP(w, r)
		}
	}
}
