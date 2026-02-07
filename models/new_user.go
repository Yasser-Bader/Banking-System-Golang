package models

import (
	"errors"
	"strconv"
)

//type User models.User

func New(firstName, lastName string, accountNumber int64, balance float64) (*User, error) {

	if firstName == "" || lastName == "" || accountNumber == 0 || balance == 0.0 {
		return nil, errors.New("the FirstName and LastName are requird")
	} else if len(strconv.Itoa(int(accountNumber))) < 14 {
		return nil, errors.New("the Account Number is requird and 14 numbers")
	} else if balance <= 100 {
		return nil, errors.New("the minimum account opening requirement is not met(minimum 100)")
	}

	return &User{
		FirstName:     firstName,
		LastName:      lastName,
		AccountNumber: accountNumber,
		Balance:       balance,
	}, nil

}
