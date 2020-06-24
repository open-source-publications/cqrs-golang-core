package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//EventStormingModel type
type EventStormingModel struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Title       string        `bson:"title" json:"title"`
	Description string        `bson:"description" json:"description"`
	CreatedAt   time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at" json:"updated_at"`
	CreatedBy   string        `bson:"created_by" json:"created_by"`
	UpdatedBY   string        `bson:"updated_by" json:"updated_by"`
}
