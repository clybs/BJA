package utils

import "github.com/globalsign/mgo"

type DB struct {
	IceCreamCollection string
	Name               string
	UserCollection     string
}

func NewDB() *DB {
	return &DB{"icecreams", "bja", "users"}
}

func (db *DB) GetSession() *mgo.Session {
	// Connect to our local mongo
	session, err := mgo.Dial("mongodb://admin1:admin1@ds121950.mlab.com:21950/bja")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	return session
}
