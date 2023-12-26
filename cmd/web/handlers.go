package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
)

type Gruz struct {
	Id    int
	Name  string
	ETSNG string
	GNG   string
}

func getGruz(db *sql.DB) []Gruz {
	gruz, err := db.Query("Select * from gruz")
	if err != nil {
		panic(err)
	}

	rez := make([]Gruz, 0)

	for gruz.Next() {
		var a Gruz
		err = gruz.Scan(&a.Id, &a.Name, &a.ETSNG, &a.GNG)
		if err != nil {
			panic(err)
		}
		rez = append(rez, a)
	}
	return rez
}

type Consignee struct {
	Id         int
	Name       string
	OKPO       string
	Was_sender bool
}

func getConsignee(db *sql.DB) []Consignee {
	gruz, err := db.Query("Select * from сonsignee")
	if err != nil {
		panic(err)
	}

	rez := make([]Consignee, 0)

	for gruz.Next() {
		var a Consignee
		err = gruz.Scan(&a.Id, &a.Name, &a.OKPO, &a.Was_sender)
		if err != nil {
			panic(err)
		}
		rez = append(rez, a)
	}
	return rez
}

type Region struct {
	Id   int
	Name string
}

func getRegion(db *sql.DB) []Region {
	gruz, err := db.Query("Select * from region")
	if err != nil {
		panic(err)
	}

	rez := make([]Region, 0)

	for gruz.Next() {
		var a Region
		err = gruz.Scan(&a.Id, &a.Name)
		if err != nil {
			panic(err)
		}
		rez = append(rez, a)
	}
	return rez
}

type Road struct {
	Id     int
	Name   string
	Length int
}

func getRoad(db *sql.DB) []Road {
	gruz, err := db.Query("Select * from road")
	if err != nil {
		panic(err)
	}

	rez := make([]Road, 0)

	for gruz.Next() {
		var a Road
		err = gruz.Scan(&a.Id, &a.Name, &a.Length)
		if err != nil {
			panic(err)
		}
		rez = append(rez, a)
	}
	return rez
}

type State struct {
	Id   int
	Name string
}

func getState(db *sql.DB) []State {
	gruz, err := db.Query("Select * from state")
	if err != nil {
		panic(err)
	}

	rez := make([]State, 0)

	for gruz.Next() {
		var a State
		err = gruz.Scan(&a.Id, &a.Name)
		if err != nil {
			panic(err)
		}
		rez = append(rez, a)
	}
	return rez
}

type Station struct {
	Id   int
	Name string
	ESR  string
	Road string
}

func getStation(db *sql.DB) []Station {
	gruz, err := db.Query("Select * from station")
	if err != nil {
		panic(err)
	}

	rez := make([]Station, 0)

	for gruz.Next() {
		var a Station
		err = gruz.Scan(&a.Id, &a.Name, &a.ESR, &a.Road)
		if err != nil {
			panic(err)
		}
		rez = append(rez, a)
	}
	return rez
}

type Wagon struct {
	Id   int
	Name string
}

func getWagon(db *sql.DB) []Wagon {
	gruz, err := db.Query("Select * from wagon")
	if err != nil {
		panic(err)
	}

	rez := make([]Wagon, 0)

	for gruz.Next() {
		var a Wagon
		err = gruz.Scan(&a.Id, &a.Name)
		if err != nil {
			panic(err)
		}
		rez = append(rez, a)
	}
	return rez
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w) // Использование помощника notFound()
		return
	}

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
	}
}

/*----------------------------------------------------------------------------------------*/

func (app *application) input(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/input.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	db, errsql := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lab3")
	if errsql != nil {
		panic(errsql)
	}

	defer db.Close()

	Goods := getGruz(db)

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
	}

	//errExTemp := ts.ExecuteTemplate(w, "app.input", Goods)
	//if errExTemp != nil {
	//	app.serverError(w, errExTemp)
	//}

	type Jopa struct {
		Goods []Gruz
	}

	err = ts.Execute(w, Jopa{Goods: Goods})
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
	}
}

/*----------------------------------------------------------------------------------------*/

func (app *application) save_article(w http.ResponseWriter, r *http.Request) {
	Number := r.FormValue("Number")
	Reg_date := r.FormValue("Reg_date")
	Status := r.FormValue("Status")
	Provide_date := r.FormValue("Provide_date")
	Departure_type := r.FormValue("Departure_type")
	Goods := r.FormValue("Goods")
	Origin_state := r.FormValue("Origin_state")
	Enter_station := r.FormValue("Enter_station")
	Region_depart := r.FormValue("Region_depart")
	Road_depart := r.FormValue("Road_depart")
	Station_depart := r.FormValue("Station_depart")
	Consigner := r.FormValue("Consigner")
	State_destination := r.FormValue("State_destination")
	Exit_station := r.FormValue("Exit_station")
	Region_destination := r.FormValue("Region_destination")
	Road_destination := r.FormValue("Road_destination")
	Station_destination := r.FormValue("Station_destination")
	Consignee := r.FormValue("Consignee")
	Wagon_type := r.FormValue("Wagon_type")
	Property := r.FormValue("Property")
	Wagon_owner := r.FormValue("Wagon_owner")
	Payer := r.FormValue("Payer")
	Road_owner := r.FormValue("Road_owner")
	Transport_manager := r.FormValue("Transport_manager")
	Tons_declared := r.FormValue("Tons_declared")
	Wagon_declared := r.FormValue("Wagon_declared")
	Wagon_accepted := r.FormValue("Wagon_accepted")
	Filing_date := r.FormValue("Filing_date")
	Agreement_date := r.FormValue("Agreement_date")
	Approval_date := r.FormValue("Approval_date")
	Start_date := r.FormValue("Start_date")
	Stop_date := r.FormValue("Stop_date")

	db, errsql := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lab3")
	if errsql != nil {
		panic(errsql)
	}

	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO 'application' "+
		"('Number', 'Reg_date', 'Status', 'Provide_date', 'Departure_type', 'Goods', 'Origin_state', "+
		"'Enter_station', 'Region_depart', 'Road_depart', 'Station_depart', 'Consigner', 'State_destination', "+
		"'Exit_station', 'Region_destination', 'Road_destination', 'Station_destination', 'Consignee', "+
		"'Wagon_type', 'Property', 'Wagon_owner', 'Payer', 'Road_owner', 'Transport_manager', 'Tons_declared', "+
		"'Tons_accepted', 'Wagon_declared', 'Wagon_accepted', 'Filing_date', 'Agreement_date', 'Approval_date', "+
		"'Start_date', 'Stop_date') VALUES('%d', '%s', '%s', '%s', '%s', '%d', '%d', '%d', '%d', '%d', '%d', "+
		"'%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%s', '%s', '%s', '%s', '%s', '%d', '%d', '%d', '%d', "+
		"'%s', '%s', '%s', '%s', '%s')"), Number, Reg_date, Status, Provide_date, Departure_type, Goods, Origin_state,
		Enter_station, Region_depart, Road_depart, Station_depart, Consigner, State_destination, Exit_station,
		Region_destination, Road_destination, Station_destination, Consignee, Wagon_type, Property, Wagon_owner,
		Payer, Road_owner, Transport_manager, Tons_declared, Wagon_declared, Wagon_accepted, Filing_date, Agreement_date,
		Approval_date, Start_date, Stop_date)
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
