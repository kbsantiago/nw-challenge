package dao

import (
	"testing"
	dao "github.com/kbsantiago/nw-challenge/src/dao"
	util "github.com/kbsantiago/nw-challenge/src/util"
 	model "github.com/kbsantiago/nw-challenge/src/model"
	"reflect"
	"log"
)

const path = "./dataset/"

var bd = dao.CompanyDAO{}

func TestGetAllCompanies(t *testing.T) {
	companies, _ := bd.GetAll()
	
	if companies == nil {
		t.Errorf("FAIL")
	}
}

func TestFindByNameAndZipcodeWhenDataMatch(t *testing.T) {
	name := "tola"
	zipcode := "78229"

	company, _ := bd.FindByNameAndZipcode(name, zipcode)

	if reflect.TypeOf(company) == reflect.TypeOf(&model.Company{}) {
		t.Errorf("FAIL")
	}
}

func TestFindByNameAndZipcodeWhenDataNotMatch(t *testing.T) {
	name := "tola"
	zipcode := "55555"
	want := ""

	company, _ := bd.FindByNameAndZipcode(name, zipcode)

	if company.CompanyName != want {
		t.Errorf("FAIL")
	}
}

func TestDataJoinerToUpdateData(t *testing.T) {
	companies, _ := bd.GetAll()
	chDataToJoin := make(chan string)
	file := "q2_clientData.csv"
	absolutePathCsv := path + file
	log.Print(path)
	go util.CsvReader(absolutePathCsv, true, chDataToJoin)
	companiesResult := bd.Joiner(chDataToJoin, companies)

	if companiesResult != nil {
		t.Errorf("FAIL")
	}

}