package dao

import (
	"strings"

	config "github.com/kbsantiago/nw-challenge/src/config"
	model "github.com/kbsantiago/nw-challenge/src/model"
	util "github.com/kbsantiago/nw-challenge/src/util"
)

var bd = CompanyDAO{}
var con = config.Config{}

func init() {
	con.Read()
	bd.Server = con.Server
	bd.Database = con.Database
	bd.Connect()
}

//Startup initialize collection
func Startup() {
	initialDataset := con.InitialDataset
	countDocuments, _ := bd.Count()

	if countDocuments == 0 {
		in := make(chan string)
		go util.CsvReader(initialDataset, true, in)

		for got := range in {
			data := strings.SplitN(got, ";", 2)
			company := model.Company{
				CompanyName: data[0],
				ZipCode:     data[1],
			}

			bd.Create(&company)
		}
	}
}
