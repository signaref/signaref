package models

import "encoding/json"

type Message struct {
	MSG     string `json:"message"`
	Success bool   `json:"success"`
	Status  int    `json:"status"`
	Trace   string `json:"trace"`
}

func (m Message) String() (s string) {
	buf, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(buf)
}

func (m Message) Bytes() []byte {
	buf, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	return buf
}
