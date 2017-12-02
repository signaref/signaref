package models

import (
	"encoding/json"
)

type Stringer interface {
	String() string
}

type Log struct {
	TimeStamp string `json:"timestamp" bson:"timestamp"`
	Status    string `json:"status" bson:"status"`
	User      string `json:"user" bson:"user"`
	Message   string `json:"message" bson:"message"`
	Source    string `json:"source" bson:"source"`
}

func (l Log) String() (s string) {
	buf, err := json.Marshal(l)
	if err != nil {
		return ""
	}
	return string(buf)
}

func (l Log) Bytes() []byte {
	buf, err := json.Marshal(l)
	if err != nil {
		return nil
	}
	return buf
}
