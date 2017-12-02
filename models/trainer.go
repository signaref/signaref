package models

import "gopkg.in/mgo.v2/bson"

type Trainer struct {
	ID        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	DOB       string        `json:"dob" bson:"dob"`
	Gender    string        `json:"gender" bson:"gender"`
	Address   string        `json:"address" bson:"address"`
	Email     string        `json:"email" bson:"email"`
	Web       string        `json:"web" bson:"web"`
	Contact   string        `json:"contact" bson:"contact"`
	Skills    []string      `json:"skills" bson:"skills"`
	Status    string        `json:"status" bson:"status"`
	TimeStamp string        `json:"timestamp" bson:"timestamp"`
}
