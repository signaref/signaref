package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	UserName  string        `json:"username" bson:"username"`
	Password  string        `json:"password" bson:"password"`
	DOB       string        `json:"dob" bson:"dob"`
	Gender    string        `json:"gender" bson:"gender"`
	Address   string        `json:"address" bson:"address"`
	Email     string        `json:"email" bson:"email"`
	Web       string        `json:"web" bson:"web"`
	Contact   string        `json:"contact" bson:"contact"`
	Contact2  string        `json:"contact2" bson:"contact2"`
	Roles     []string      `json:"roles" bson:"roles"`
	Status    string        `json:"status" bson:"status"`
	TimeStamp string        `json:"timestamp" bson:"timestamp"`
}

func ValidateUser(u User) string {
	if u.UserName == "" {
		return "User Name field is empty"
	}
	if u.Password == "" {
		return "Password field is empty"
	}
	if u.DOB == "" {
		return "Date of birth field is empty"
	}
	if u.Gender == "" {
		return "Gender field is empty"
	}
	if u.Email == "" {
		return "Email field is empty"
	}

	if u.Contact == "" {
		return "Contact number field is empty"
	}
	if len(u.Roles) < 1 {
		return "Roles field is empty"
	}

	if u.Status == "" {
		return "Status field is empty"
	}
	if u.TimeStamp == "" {
		return "Timestamp is empty"
	}
	return ""
}

func ValidateUserForUpdate(u User) string {
	if u.ID.String() == "" {
		return "id field is empty"
	}
	return ""
}
