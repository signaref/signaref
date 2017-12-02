package models

import "gopkg.in/mgo.v2/bson"

type Map struct {
	Lat string `json:"lat" bson:"lat"`
	Lng string `json:"lng" bson:"lng"`
}

type Venue struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	Address   string        `json:"address" bson:"address"`
	Location  Map           `json:"location" bson:"location"`
	Email     string        `json:"email" bson:"email"`
	Web       string        `json:"web" bson:"web"`
	Contact   string        `json:"contact" bson:"contact"`
	Status    string        `json:"status" bson:"status"`
	TimeStamp string        `json:"timestamp" bson:"timestamp"`
}
