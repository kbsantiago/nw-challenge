package domain

import (
	"gopkg.in/mgo.v2/bson"
)

//Company define company attributes
type Company struct {
	ID          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CompanyName string        `bson:"name" json:"name"`
	ZipCode     string        `bson:"zip" json:"zip"`
	Website     string        `bson:"website" json:"website"`
}
