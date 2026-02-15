package controllers

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func GetUserInput(promptText string) (string, error) {
	fmt.Printf(" %v", promptText)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return "", errors.New("invalid input ")

	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	if text==""{
		return "",errors.New("input cannot be empty")
	}
	return text, nil

}
