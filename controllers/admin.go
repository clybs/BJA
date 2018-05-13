package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/clybs/BJA/models"
	"github.com/clybs/BJA/utils"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gocraft/web"
	"net/http"
)

type AdminController struct {
	session *mgo.Session
	db      *utils.DB
}

var ut utils.Auth

// Add session to our AdminController
func NewAdminController(session *mgo.Session) *AdminController {
	return &AdminController{session, utils.NewDB()}
}

func (ac *AdminController) Login(rw web.ResponseWriter, req *web.Request) {
	// Get the username
	username := req.FormValue("username")

	// Get the password
	password := req.FormValue("password")

	// Initialize the model to use
	modelU := models.User{}

	// Hash the password
	hash, err := ut.HashPassword(password)
	if err != nil {
		http.Error(rw, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Create the query
	query := bson.M{"username": username}

	// Try searching for the username
	if err := ac.session.DB(ac.db.Name).C(ac.db.UserCollection).Find(query).One(&modelU); err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	// Check if hash and password matches
	match := ut.CheckPasswordHash(password, hash)
	if !match {
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Create the token
	token := ut.CreateToken(modelU.Id.String())
	mToken := map[string]string{"token": token}
	mTokenM, _ := json.Marshal(mToken)

	// Display results
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(mTokenM))
}
