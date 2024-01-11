package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	m := http.NewServeMux()
	m.HandleFunc("/", app.home)
	m.HandleFunc("/input", app.input)
	m.HandleFunc("/save_application", app.save_application)
	m.HandleFunc("/add_gruz", app.addGruz)
	m.HandleFunc("/save_gruz", app.save_gruz)
	m.HandleFunc("/add_state", app.add_state)
	m.HandleFunc("/save_state", app.save_state)
	m.HandleFunc("/add_station", app.add_station)
	m.HandleFunc("/save_station", app.save_station)
	m.HandleFunc("/add_region", app.add_region)
	m.HandleFunc("/save_region", app.save_region)
	m.HandleFunc("/add_road", app.add_road)
	m.HandleFunc("/save_road", app.save_road)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	m.Handle("/static/", http.StripPrefix("/static", fileServer))

	return m
}
