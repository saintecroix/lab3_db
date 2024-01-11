package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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
	Was_sender int
}

func getConsignee(db *sql.DB) []Consignee {
	gruz, err := db.Query("Select * from consignee")
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

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	type Jopa struct {
		Gruz      []Gruz
		Consignee []Consignee
		Region    []Region
		Road      []Road
		State     []State
		Station   []Station
		Wagon     []Wagon
	}

	err = ts.Execute(w, Jopa{Gruz: getGruz(db), Consignee: getConsignee(db), Region: getRegion(db), Road: getRoad(db),
		State: getState(db), Station: getStation(db), Wagon: getWagon(db)})
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

/*----------------------------------------------------------------------------------------*/

func (app *application) save_application(w http.ResponseWriter, r *http.Request) {
	number, _ := strconv.Atoi(r.FormValue("Number"))
	reg_date := r.FormValue("Reg_date")
	status := r.FormValue("Status")
	provide_date := r.FormValue("Provide_date")
	departure_type := r.FormValue("Departure_type")
	property := r.FormValue("Property")
	wagon_owner := r.FormValue("Wagon_owner")
	payer := r.FormValue("Payer")
	road_owner := r.FormValue("Road_owner")
	transport_manager := r.FormValue("Transport_manager")
	tons_declared, _ := strconv.Atoi(r.FormValue("Tons_declared"))
	tons_accepted, _ := strconv.Atoi(r.FormValue("Tons_accepted"))
	wagon_declared, _ := strconv.Atoi(r.FormValue("Wagon_declared"))
	wagon_accepted, _ := strconv.Atoi(r.FormValue("Wagon_accepted"))
	filing_date := r.FormValue("Filing_date")
	agreement_date := r.FormValue("Agreement_date")
	approval_date := r.FormValue("Approval_date")
	start_date := r.FormValue("Start_date")
	stop_date := r.FormValue("Stop_date")

	db, errsql := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lab3")
	if errsql != nil {
		panic(errsql)
	}

	defer db.Close()

	goods, _ := strconv.Atoi(r.FormValue("Goods"))

	origin_state, _ := strconv.Atoi(r.FormValue("Origin_state"))

	enter_station, _ := strconv.Atoi(r.FormValue("Enter_station"))

	region_depart, _ := strconv.Atoi(r.FormValue("Region_depart"))

	road_depart, _ := strconv.Atoi(r.FormValue("Road_depart"))

	station_depart, _ := strconv.Atoi(r.FormValue("Station_depart"))

	consigner, _ := strconv.Atoi(r.FormValue("Consigner"))

	state_destination, _ := strconv.Atoi(r.FormValue("State_destination"))

	exit_station, _ := strconv.Atoi(r.FormValue("Exit_station"))

	region_destination, _ := strconv.Atoi(r.FormValue("Region_destination"))

	road_destination, _ := strconv.Atoi(r.FormValue("Road_destination"))

	station_destination, _ := strconv.Atoi(r.FormValue("Station_destination"))

	consignees, _ := strconv.Atoi(r.FormValue("Consignee"))

	wagon_type, _ := strconv.Atoi(r.FormValue("Wagon_type"))

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `application` "+
		"(`Number`, `Reg_date`, `Status`, `Provide_date`, `Departure_type`, `Goods`, `Origin_state`, "+
		"`Enter_station`, `Region_depart`, `Road_depart`, `Station_depart`, `Consigner`, `State_destination`, "+
		"`Exit_station`, `Region_destination`, `Road_destination`, `Station_destination`, `Consignee`, "+
		"`Wagon_type`, `Property`, `Wagon_owner`, `Payer`, `Road_owner`, `Transport_manager`, `Tons_declared`, "+
		"`Tons_accepted`, `Wagon_declared`, `Wagon_accepted`, `Filing_date`, `Agreement_date`, `Approval_date`, "+
		"`Start_date`, `Stop_date`) VALUES('%d', '%s', '%s', '%s', '%s', '%d', '%d', '%d', '%d', '%d', '%d', "+
		"'%d', '%d', '%d', '%d', '%d', '%d', '%d', '%d', '%s', '%s', '%s', '%s', '%s', '%d', '%d', '%d', '%d', "+
		"'%s', '%s', '%s', '%s', '%s')", number, reg_date, status, provide_date, departure_type, goods, origin_state,
		enter_station, region_depart, road_depart, station_depart, consigner, state_destination, exit_station,
		region_destination, road_destination, station_destination, consignees, wagon_type, property, wagon_owner,
		payer, road_owner, transport_manager, tons_declared, tons_accepted, wagon_declared, wagon_accepted, filing_date,
		agreement_date, approval_date, start_date, stop_date))
	if err != nil {
		app.serverError(w, err)
		return
	}
	defer insert.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

/*----------------------------------------------------------------------------------------*/

func (app *application) addGruz(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/add_gruz.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	db, errsql := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lab3")
	if errsql != nil {
		panic(errsql)
	}

	defer db.Close()

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	type Jopa struct {
		Gruz []Gruz
	}

	err = ts.Execute(w, Jopa{Gruz: getGruz(db)})
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

/*----------------------------------------------------------------------------------------*/

func (app *application) save_gruz(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")
	etsng := r.FormValue("ETSNG")
	gng := r.FormValue("GNG")

	db, errsql := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lab3")
	if errsql != nil {
		panic(errsql)
	}

	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `gruz` "+
		"(`Name`, `ETSNG`, `GNG`) VALUES('%s', '%s', '%s')", name, etsng, gng))

	if err != nil {
		app.serverError(w, err)
		return
	}
	defer insert.Close()

	http.Redirect(w, r, "/input", http.StatusSeeOther)
}

/*----------------------------------------------------------------------------------------*/

func (app *application) add_state(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/add_state.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	db, errsql := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lab3")
	if errsql != nil {
		panic(errsql)
	}

	defer db.Close()

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	type Jopa struct {
		State []State
	}

	err = ts.Execute(w, Jopa{State: getState(db)})
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

/*----------------------------------------------------------------------------------------*/

func (app *application) save_state(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")

	db, errsql := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lab3")
	if errsql != nil {
		panic(errsql)
	}

	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `state` "+
		"(`Name`) VALUES('%s')", name))

	if err != nil {
		app.serverError(w, err)
		return
	}
	defer insert.Close()

	http.Redirect(w, r, "/input", http.StatusSeeOther)
}

/*----------------------------------------------------------------------------------------*/

func (app *application) add_station(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/add_station.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	db, errsql := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lab3")
	if errsql != nil {
		panic(errsql)
	}

	defer db.Close()

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	type Jopa struct {
		Station []Station
	}

	err = ts.Execute(w, Jopa{Station: getStation(db)})
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

/*----------------------------------------------------------------------------------------*/

func (app *application) save_station(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")
	esr := r.FormValue("ESR")
	road := r.FormValue("Road")

	db, errsql := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lab3")
	if errsql != nil {
		panic(errsql)
	}

	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `station` "+
		"(`Name`,`ESR`, `Road`) VALUES('%s', `%s`, `%s`)", name, esr, road))

	if err != nil {
		app.serverError(w, err)
		return
	}
	defer insert.Close()

	http.Redirect(w, r, "/input", http.StatusSeeOther)
}

/*----------------------------------------------------------------------------------------*/

func (app *application) add_region(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/add_region.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	db, errsql := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lab3")
	if errsql != nil {
		panic(errsql)
	}

	defer db.Close()

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	type Jopa struct {
		Region []Region
	}

	err = ts.Execute(w, Jopa{Region: getRegion(db)})
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

/*----------------------------------------------------------------------------------------*/

func (app *application) save_region(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")

	db, errsql := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lab3")
	if errsql != nil {
		panic(errsql)
	}

	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `region` "+
		"(`Name`) VALUES('%s')", name))

	if err != nil {
		app.serverError(w, err)
		return
	}
	defer insert.Close()

	http.Redirect(w, r, "/input", http.StatusSeeOther)
}

/*----------------------------------------------------------------------------------------*/

func (app *application) add_road(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/add_road.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	db, errsql := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lab3")
	if errsql != nil {
		panic(errsql)
	}

	defer db.Close()

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}

	type Jopa struct {
		Road []Road
	}

	err = ts.Execute(w, Jopa{Road: getRoad(db)})
	if err != nil {
		app.serverError(w, err) // Использование помощника serverError()
		return
	}
}

/*----------------------------------------------------------------------------------------*/

func (app *application) save_road(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("Name")
	length, _ := strconv.Atoi(r.FormValue("Length"))

	db, errsql := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/lab3")
	if errsql != nil {
		panic(errsql)
	}

	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `road` "+
		"(`Name`, `Length`) VALUES('%s', '%d')", name, length))

	if err != nil {
		app.serverError(w, err)
		return
	}
	defer insert.Close()

	http.Redirect(w, r, "/input", http.StatusSeeOther)
}

//func getid(db *sql.DB, col string, name string) int {
//	raw, err := db.Query(fmt.Sprintf("SELECT id from %s where Name = %s", col, name))
//	if err != nil {
//		panic(err)
//	}
//	defer raw.Close()
//	var res int
//	errrod := raw.Scan(&res)
//	if errrod != nil {
//		panic(errrod)
//	}
//	return res
//}
