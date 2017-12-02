package models

import "gopkg.in/mgo.v2/bson"

type Role struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	Desc      string        `json:"desc" bson:"desc"`
	Status    string        `json:"status" bson:"status"`
	TimeStamp string        `json:"timestamp" bson:"timestamp"`
}

func ValidateRole(r Role) string {
	if r.Name == "" {
		return "Role Name field is empty"
	}
	if r.Desc == "" {
		return "Description field is empty"
	}

	if r.Status == "" {
		return "Status field is empty"
	}
	if r.TimeStamp == "" {
		return "Timestamp is empty"
	}
	return ""
}

func ValidateRoleForUpdate(r Role) string {
	if r.ID.String() == "" {
		return "id field is empty"
	}
	return ""
}
