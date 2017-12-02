// models project models.go
package models

import "gopkg.in/mgo.v2/bson"

// service type model
type Service_Type struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Service     string        `json:"service" bson:"service"`
	Description string        `json:"desc" bson:"desc"`
	Image       string        `json:"image" bson:"image"`
	Status      string        `json:"status" bson:"status"`
	Catalog     []string      `json:"catalog" bson:"catalog"`
	Timestamp   string        `json:"timestamp" bson:"timestamp"`
}

// Validating each and every field of Service_Type object.
// Any additional validations , can be developed here..
func ValidateService_Type(st Service_Type) string {
	if st.Service == "" {
		return "Service field is empty"
	}
	if st.Description == "" {
		return "Description field is empty"
	}
	if st.Image == "" {
		return "Image field is empty"
	}
	if st.Status == "" {
		return "Status field is empty"
	}
	if len(st.Catalog) <= 0 {
		return "Catalog is empty"
	}
	if st.Timestamp == "" {
		return "Timestamp is empty"
	}
	return ""
}

func ValidateService_TypeforUpdate(st Service_Type) string {
	if st.ID.String() == "" {
		return "id field is empty"
	}
	return ""
}

type Service_Provider struct {
	ID          bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Service     string        `json:"service" bson:"service"`
	Description string        `json:"desc" bson:"desc"`
	Image       string        `json:"image" bson:"image"`
	Status      string        `json:"status" bson:"status"`
	Catalog     []string      `json:"catalog" bson:"catalog"`
	Timestamp   string        `json:"timestamp" bson:"timestamp"`
}

type Vendor struct {
	ID                   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Company              string
	Description          string
	Address              string
	Email                string
	Web                  string
	Social_Media         []string
	Contact_Person       string
	Contact_Person_email string
	Status               string
	TimeStamp            string
}
