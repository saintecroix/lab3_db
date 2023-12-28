package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	m := http.NewServeMux()
	r := mux.NewRouter()
	r.HandleFunc("/", app.home).Methods("GET")
	r.HandleFunc("/input", app.input).Methods("GET")
	r.HandleFunc("/save_application", app.save_application).Methods("POST")

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	m.Handle("/static/", http.StripPrefix("/static", fileServer))

	return m
}
