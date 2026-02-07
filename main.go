package main

import (
	"fmt"
	"strconv"

	"github.com/Banking-System/controllers"
	"github.com/Banking-System/models"
)

func main() {
	fmt.Println("Welcome to the Banking System")

	fmt.Println("What do you want to do?")
	fmt.Println("1. Login")
	fmt.Println("2. New Account")
	fmt.Println("3. Account Detilse")
	fmt.Println("4. Check Balance")
	fmt.Println("5. Deposit Money")
	fmt.Println("6. Withdraw Money")
	fmt.Println("7. Exit")
	var choice int
	fmt.Print("Your Choice:")
	fmt.Scan(&choice)
	switch choice {
	case 2:
		userFirstName, err := controllers.GetUserInput("Place Enter Your First Name :")
		if err != nil {
			fmt.Println(err)
		}
		userLastName, err := controllers.GetUserInput("Place Enter Your Last Name :")
		if err != nil {
			fmt.Println(err)
		}
		userAccountNum, err := controllers.GetUserInput("Place Enter Your Account Number :")
		if err != nil {
			fmt.Println(err)
		}
		userBalance, err := controllers.GetUserInput("Place Enter Your Balance :")
		if err != nil {
			fmt.Println(err)
		}

		//firstName, _ := userFirstName.(string)
		//lastName, _ := userLastName.(string)
		accountNum, err := strconv.ParseInt(userAccountNum, 10, 64)
		balance, err := strconv.ParseFloat(userBalance, 64)

		var appUser *models.User

		appUser, err = models.New(userFirstName, userLastName, accountNum, balance)
		if err != nil {
			fmt.Println(err)
		}
		//appUser.OutputUserDetails()
		appUser.SaveAccount()
	case 3:
		data, err := models.GetAccount()
		if err != nil {
			fmt.Println(err)
			fmt.Println("Don't have account")

		}
		fmt.Printf("first name : %v \n", data.FirstName)
		fmt.Printf("last name : %v \n", data.LastName)
		fmt.Printf("Account number : %v \n", data.AccountNumber)
		fmt.Printf("Balance %.2f", data.Balance)
	default:
		fmt.Println("good bay")
	}

}
