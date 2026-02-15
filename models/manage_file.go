package models

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
)

//const accountFile = "accountFile.json"

func (u *User) SaveAccount() error {

	accountInfo := strings.ReplaceAll(u.AccountNumber, " ", "_") + ".json"
	jsonBytes, err := json.MarshalIndent(u, "", " ")
	if err != nil {
		return err
	}
	os.WriteFile(accountInfo, jsonBytes, 0644)
	return nil
}

func GetAccount(accountNumber string) (*User, error) {

	fileName := strings.ReplaceAll(accountNumber, " ", "_") + ".json"
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var myData User
	err = json.Unmarshal(fileContent, &myData)
	if err != nil {
		return nil, errors.New("Error parsing JSON")
	}

	return &myData, nil

}
