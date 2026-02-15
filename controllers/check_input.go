package controllers

import (
	"fmt"
	"strconv"

	"github.com/Banking-System/models"
)

func CheckInput() {

	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Login")
		fmt.Println("2. New Account")
		fmt.Println("3. Account Details")
		fmt.Println("4. Check Balance")
		fmt.Println("5. Deposit Money")
		fmt.Println("6. Withdraw Money")
		fmt.Println("7. Exit")

		choiceString, err := GetUserInput("Your Choice:")
		if err != nil {
			fmt.Println("invalid input")
			continue
		}
		choice, err := strconv.Atoi(choiceString)
		if err != nil {
			fmt.Println("Please enter a valid number ")
			continue
		}
		switch choice {
		case 1:
			accountNumber, err := GetUserInput("Enter Your Account number: ")
			data, err := models.GetAccount(accountNumber)
			if err != nil {
				fmt.Println("Don't have account")
				continue

			}

			fmt.Printf("Welcome %v \n", data.FirstName)
			fmt.Printf("Your Account number is: %v \n", data.AccountNumber)
			fmt.Printf("Your Balance is: %.2f \n\n", data.Balance)
			continue

		case 2:
			userFirstName, err := GetUserInput("Place Enter Your First Name :")
			if err != nil {
				fmt.Println(err)
				continue
			}
			userLastName, err := GetUserInput("Place Enter Your Last Name :")
			if err != nil {
				fmt.Println(err)
				continue
			}
			userAccountNum, err := GetUserInput("Place Enter Your Account Number :")
			if err != nil {
				fmt.Println(err)
				continue
			}
			userBalance, err := GetUserInput("Place Enter Your Balance :")
			if err != nil {
				fmt.Println(err)
				continue
			}

			balance, err := strconv.ParseFloat(userBalance, 64)

			var appUser *models.User

			appUser, err = models.New(userFirstName, userLastName, userAccountNum, balance)
			if err != nil {
				fmt.Println(err)
				continue
			}
			err = appUser.SaveAccount()
			if err != nil {
				fmt.Println("error: don't save file ")
				continue
			}
			fmt.Println("saved file success")

		case 3:
			fmt.Print("sssss")
			continue
		case 4:
			accountNumber, err := GetUserInput("Enter Your Account number: ")
			data, err := models.GetAccount(accountNumber)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Don't have account")
				continue
			}
			fmt.Printf("Hi %v \n", data.FirstName)
			fmt.Printf("Your Balance is: %.2f \n \n", data.Balance)
			continue
		case 5:
			accountNumber, err := GetUserInput("Enter Your Account number: ")
			data, err := models.GetAccount(accountNumber)
			if err != nil {
				fmt.Println("Don't have account")
				continue

			}
			amountStr, err := GetUserInput("Enter Amount: ")
			if err != nil {
				fmt.Println("invalid Amount input")
				continue

			}
			amountFlo, err := strconv.ParseFloat(amountStr, 64)
			if amountFlo < 0 {
				fmt.Println("the amount is less than zero")
				continue
			}

			userData, err := data.Deposit(amountFlo)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Deposited success \n Account Name: %v And Account number is: %v ", userData.FirstName, userData.AccountNumber)
			fmt.Printf("Your New Balance is: %.2f \n\n", userData.Balance)
		case 6:
			accountNumber, err := GetUserInput("Enter Your Account number: ")
			data, err := models.GetAccount(accountNumber)
			if err != nil {
				fmt.Println("Don't have account")

			}
			amountStr, _ := GetUserInput("Enter Amount: ")

			amountFlo, err := strconv.ParseFloat(amountStr, 64)
			if data.Balance < amountFlo {
				fmt.Println("Your Balance is insufficient")
				continue

			}
			if err != nil {
				fmt.Println("invalid Amount input")
				continue

			}
			userData, err := data.Withdrawal(amountFlo)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Withered success \n Account Name: %v And Account number is: %v ", userData.FirstName, userData.AccountNumber)
			fmt.Printf("Your New Balance is: %.2f \n\n", userData.Balance)
		default:
			fmt.Println("Goodbye")
			return
		}
	}

}
