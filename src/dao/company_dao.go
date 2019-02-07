package dao

import (
	"log"
	"strings"

	model "github.com/kbsantiago/nw-challenge/src/model"
	util "github.com/kbsantiago/nw-challenge/src/util"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//CompanyDAO struct to database connection
type CompanyDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

func collection() *mgo.Collection {
	return db.C("company")
}

//Connect manager database connection
func (c *CompanyDAO) Connect() {
	session, err := mgo.Dial(c.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(c.Database)
}

//GetAll saves company information in database
func (c *CompanyDAO) GetAll() ([]model.Company, error) {
	var companies []model.Company
	err := collection().Find(bson.M{}).All(&companies)
	return companies, err
}

//FindByNameAndZipcode saves company information in database
func (c *CompanyDAO) FindByNameAndZipcode(name string, zipcode string) (model.Company, error) {
	var company model.Company
	conditions := bson.M{"name": bson.M{"$regex": bson.RegEx{`.*` + name + `.*`, ""}},
		"$and": []bson.M{bson.M{"zip": bson.M{"$eq": zipcode}}}}
	err := collection().Find(conditions).One(&company)
	return company, err
}

//Create saves company information in database
func (c *CompanyDAO) Create(company *model.Company) error {
	err := collection().Insert(&company)
	return err
}

//Update saves company information in database
func (c *CompanyDAO) Update(company *model.Company) error {
	err := collection().Update(bson.M{"name": company.CompanyName, "zip": company.ZipCode}, &company)
	return err
}

//Count return the total number of documents
func (c *CompanyDAO) Count() (int, error) {
	tdocs, err := collection().Count()
	return tdocs, err
}

//DataIntegration merge data
func (c *CompanyDAO) DataIntegration() {
	integrationDataset := con.UpdateDataset

	in := make(chan string)
	go util.CsvReader(integrationDataset, true, in)

	for got := range in {
		data := strings.SplitN(got, ";", 3)
		company := model.Company{
			CompanyName: data[0],
			ZipCode:     data[1],
			Website:     data[2],
		}
		bd.Update(&company)
	}
}

//Joiner teste
func (c *CompanyDAO) Joiner(dataToJoin chan string, companies []model.Company) error {
	var err error
	for item := range dataToJoin {
		data := strings.SplitN(item, ";", 3)
		company := model.Company{
			CompanyName: data[0],
			ZipCode:     data[1],
			Website:     data[2],
		}
		for i := range companies {
			companyName := companies[i].CompanyName
			zipCode := companies[i].ZipCode
			if companyName == company.CompanyName && zipCode == company.ZipCode {
				companies[i].Website = company.Website
				err = bd.Update(&company)
			}
		}
	}
	return err;
}
