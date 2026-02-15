package models

type User struct {
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	AccountNumber string  `json:"account_number"`
	Balance       float64 `json:"balance"`
}