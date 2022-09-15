package models

type Response struct {
	Status  int         `json:"STATUS"`
	Message string      `json:"MESSAGE"`
	Data    interface{} `json:"DATA"`
}
