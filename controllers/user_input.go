package controllers

import (
	"errors"
	"fmt"
)

// type chickValue interface
func GetUserInput(promptText string) (string, error) {
	var value string
	fmt.Print(promptText)
	fmt.Scanln(&value)
	if value == "" {
		return value, errors.New("this fild is required")
	}
	return value, nil
}
