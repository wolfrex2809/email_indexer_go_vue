package models

type CustomError struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
}
