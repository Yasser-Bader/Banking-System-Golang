package models

import (
	"errors"
)

func New(firstName, lastName string, accountNumber string, balance float64) (*User, error) {

	if firstName == "" || lastName == "" || accountNumber == "" || balance == 0.0 {
		return nil, errors.New("the FirstName and LastName are required")
	} else if len(accountNumber) < 14 {
		return nil, errors.New("the Account Number is required and 14 numbers")
	} else if balance < 100 {
		return nil, errors.New("the minimum account opening requirement is not met(minimum 100)")
	}

	return &User{
		FirstName:     firstName,
		LastName:      lastName,
		AccountNumber: accountNumber,
		Balance:       balance,
	}, nil

}
