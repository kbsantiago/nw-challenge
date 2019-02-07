package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	config "github.com/kbsantiago/nw-challenge/src/config"
	"github.com/kbsantiago/nw-challenge/src/dao"
	util "github.com/kbsantiago/nw-challenge/src/util"
	"github.com/rs/cors"

	"github.com/gorilla/mux"
)

var bd = dao.CompanyDAO{}
var con = config.Config{}

func init() {
	con.Read()
	bd.Server = con.Server
	bd.Database = con.Database
	bd.Connect()
}


//GetAll return all companies
func GetAll(w http.ResponseWriter, r *http.Request) {
	companies, _ := bd.GetAll()
	jcompanies, _ := json.Marshal(companies)
	fmt.Fprintln(w, string(jcompanies))
}

//FindByNameAndZipcode return single company based on search parameters
func FindByNameAndZipcode(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	company, _ := bd.FindByNameAndZipcode(params["name"], params["zip"])
	jcompany, _ := json.Marshal(company)
	fmt.Fprintln(w, string(jcompany))
}

//Update company
func Update(w http.ResponseWriter, r *http.Request) {
	chDataToJoin := make(chan string)
	go util.CsvReader(con.UpdateDataset, true, chDataToJoin)
	companies, _ := bd.GetAll()
	bd.Joiner(chDataToJoin, companies)
	jcompanies, _ := json.Marshal(companies)
	fmt.Fprintln(w, string(jcompanies))
}

//Server webservers
func Server() {
	router := mux.NewRouter()
	router.HandleFunc("/", GetAll).Methods("GET")
	router.HandleFunc("/company/update", Update).Methods("GET")
	router.HandleFunc("/company/{name}/{zip}", FindByNameAndZipcode).Methods("GET")

	handler := cors.Default().Handler(router)

	if err := http.ListenAndServe(":3000", handler); err != nil {
		log.Fatal(err)
	}
}

