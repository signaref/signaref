package models

import "gopkg.in/mgo.v2/bson"

type Scope struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	URI       string        `json:"uri" bson:"uri"`
	Scope     string        `json:"scope" bson:"scope"`
	Module    string        `json:"module" bson:"module"`
	Method    string        `json:"method" bson:"method"`
	Status    string        `json:"status" bson:"status"`
	TimeStamp string        `json:"timestamp" bson:"timestamp"`
}

/*func ValidateScope(s Scope) string {
	if s.Name == "" {
		return "Role Name field is empty"
	}
	if s.Desc == "" {
		return "Description field is empty"
	}

	if s.Scope == "" {
		return "scope field is empty"
	}
	if s.Scope_Desc == "" {
		return "Screen Description field is empty"
	}

	if s.Status == "" {
		return "Status field is empty"
	}
	if s.TimeStamp == "" {
		return "Timestamp is empty"
	}
	return ""
}

func ValidateScopeForUpdate(s Scope) string {
	if s.ID.String() == "" {
		return "id field is empty"
	}
	return ""
}
*/
