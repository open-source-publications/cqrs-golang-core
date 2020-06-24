package repository

import (
	"log"

	"github.com/open-source-publications/cqrs-golang-core/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//EventStormingRepository type
type EventStormingRepository struct {
	Server   string
	Database string
}

var db *mgo.Database

//COLLECTION
const (
	COLLECTION = "eventstorming"
)

//Connect function
func (e *EventStormingRepository) Connect() {
	session, err := mgo.Dial(e.Server)

	if err != nil {
		log.Fatal(err)
	}

	db = session.DB(e.Database)
}

//Save this function insert in databse
func (e *EventStormingRepository) Save(eventStormingModel model.EventStormingModel) error {
	err := db.C(COLLECTION).Insert(&eventStormingModel)
	return err
}

//Delete a eventstorming of database
func (e *EventStormingRepository) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

//Update a eventStorming in database
func (e *EventStormingRepository) Update(id string, eventStormingModel model.EventStormingModel) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &eventStormingModel)
	return err
}

//GetByID in database
func (e *EventStormingRepository) GetByID(id string) (model.EventStormingModel, error) {
	var eventStormingModel model.EventStormingModel

	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&eventStormingModel)
	return eventStormingModel, err
}

//GetAll Elements in collection
func (e *EventStormingRepository) GetAll() ([]model.EventStormingModel, error) {
	var eventStormingModel []model.EventStormingModel

	err := db.C(COLLECTION).Find(bson.M{}).All(&eventStormingModel)
	return eventStormingModel, err
}
