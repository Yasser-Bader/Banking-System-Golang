package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const accountFile = "accountFile.json"

func (u *User) SaveAccount() {

	accountInfo := strings.ReplaceAll(u.FirstName, " ", "_") + ".json"
	jsonBytes, err := json.MarshalIndent(u, "", " ")
	if err != nil {
		fmt.Println("erooor")
		return
	}
	os.WriteFile(accountInfo, jsonBytes, 0644)
	fmt.Println("Seved Account")
}

func GetAccount() (*User, error) {
	var name string
	fmt.Print("Your name: ")
	fmt.Scanln(&name)
	fileName := strings.ReplaceAll(name, " ", "_") + ".json"
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		return nil, errors.New("Error Reading file")
	}
	var myData User
	err = json.Unmarshal(fileContent, &myData)
	if err != nil {
		return nil, errors.New("Error parsing JSON")
	}

	return &myData, nil

}
