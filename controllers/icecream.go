package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/clybs/BJA/models"
	"github.com/clybs/BJA/utils"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/gocraft/web"
	"io/ioutil"
	"net/http"
)

type IceCreamController struct {
	session *mgo.Session
	db      *utils.DB
}

// Add session to our AdminController
func NewIceCreamController(session *mgo.Session) *IceCreamController {
	return &IceCreamController{session, utils.NewDB()}
}

func (ic *IceCreamController) displayResults(modelIC *models.IceCream, rw web.ResponseWriter, httpStatus int) {
	jsonIC, err := json.Marshal(modelIC)
	if err != nil {
		fmt.Println(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(httpStatus)
	fmt.Fprintf(rw, "%s\n", jsonIC)
}

func (ic *IceCreamController) Create(rw web.ResponseWriter, req *web.Request) {
	modelIC := models.IceCream{}

	// Decode json
	json.NewDecoder(req.Body).Decode(&modelIC)

	// Create bson ID
	modelIC.Id = bson.NewObjectId()

	// Store the ice cream in mongodb
	ic.session.DB(ic.db.Name).C(ic.db.IceCreamCollection).Insert(modelIC)

	// Display results
	ic.displayResults(&modelIC, rw, http.StatusCreated)
}

func (ic *IceCreamController) List(rw web.ResponseWriter, req *web.Request) {
	// Initialize the model to use
	var modelICs models.IceCreams

	// Fetch ice cream
	if err := ic.session.DB(ic.db.Name).C(ic.db.IceCreamCollection).Find(nil).All(&modelICs); err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	// Display results
	modelICsM, _ := json.Marshal(modelICs)

	// Display results
	rw.Header().Set("Content-Type", "application/json")
	fmt.Fprint(rw, string(modelICsM))
}

func (ic *IceCreamController) Read(rw web.ResponseWriter, req *web.Request) {
	// Get the id
	id := req.PathParams["id"]

	// Verify id is ObjectId hex representation, otherwise return status not found
	if !bson.IsObjectIdHex(id) {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	// ObjectIdHex returns an ObjectId from the provided hex representation.
	oId := bson.ObjectIdHex(id)

	// Initialize the model to use
	modelIC := models.IceCream{}

	// Fetch ice cream
	if err := ic.session.DB(ic.db.Name).C(ic.db.IceCreamCollection).FindId(oId).One(&modelIC); err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	// Display results
	ic.displayResults(&modelIC, rw, http.StatusOK)
}

func (ic *IceCreamController) Update(rw web.ResponseWriter, req *web.Request) {
	// Get the id
	id := req.PathParams["id"]

	// Verify id is ObjectId hex representation, otherwise return status not found
	if !bson.IsObjectIdHex(id) {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	// ObjectIdHex returns an ObjectId from the provided hex representation.
	oId := bson.ObjectIdHex(id)

	// Initialize the model to use
	modelIC := models.IceCream{}

	// Read the body
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// Fetch ice cream
	if err := ic.session.DB(ic.db.Name).C(ic.db.IceCreamCollection).FindId(oId).One(&modelIC); err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	// Transfer body to model
	err = json.Unmarshal(body, &modelIC)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// Update ice cream
	if err := ic.session.DB(ic.db.Name).C(ic.db.IceCreamCollection).UpdateId(oId, &modelIC); err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	// Display results
	ic.displayResults(&modelIC, rw, http.StatusOK)
}

func (ic *IceCreamController) Delete(rw web.ResponseWriter, req *web.Request) {
	// Get the id
	id := req.PathParams["id"]

	// Verify id is ObjectId hex representation, otherwise return status not found
	if !bson.IsObjectIdHex(id) {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	// ObjectIdHex returns an ObjectId from the provided hex representation.
	oId := bson.ObjectIdHex(id)

	// Delete user
	if err := ic.session.DB(ic.db.Name).C(ic.db.IceCreamCollection).RemoveId(oId); err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	// Display results
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, "Deleted ice cream with id: ", id, "\n")
}
