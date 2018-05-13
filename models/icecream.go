package models

import "github.com/globalsign/mgo/bson"

type IceCream struct {
	AllergyInfo           string        `json:"allergy_info" bson:"allergy_info"`
	Description           string        `json:"description" bson:"description"`
	DietaryCertifications string        `json:"dietary_certifications" bson:"dietary_certifications"`
	Id                    bson.ObjectId `json:"id" bson:"_id"`
	ImageClosed           string        `json:"image_closed" bson:"image_closed"`
	ImageOpen             string        `json:"image_open" bson:"image_open"`
	Ingredients           []string      `json:"ingredients" bson:"ingredients"`
	Name                  string        `json:"name" bson:"name"`
	ProductID             string        `json:"productId" bson:"productId"`
	SourcingValues        []string      `json:"sourcing_values" bson:"sourcing_values"`
	Story                 string        `json:"story" bson:"story"`
}

type IceCreams []IceCream
