package models

import "github.com/globalsign/mgo/bson"

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Password string        `json:"password" bson:"password"`
	Username string        `json:"username" bson:"username"`
}

type Users []User
