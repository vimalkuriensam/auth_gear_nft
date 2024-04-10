package models

type Payload struct {
	Id   string
	Kind string
	Type string
	Data []byte
}

type PayloadData struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Data      []byte `json:"data"`
	Code      int    `json:"code"`
	TimeStamp string `json:"timestamp"`
}
