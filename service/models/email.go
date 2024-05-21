package models

type Email struct {
	MessageID string `json:"message_id"`
	Date      string `json:"date"`
	From      string `json:"from"`
	To        string `json:"to"`
	Subject   string `json:"subject"`
	Content   string `json:"content"`
}

type EmailSearchResponse struct {
	Status string  `json:"status"`
	Data   []Email `json:"data"`
}
