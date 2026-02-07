package models

type User struct {
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	AccountNumber int64   `json:"account_number"`
	Balance       float64 `json:"balance"`
}

/*func (u *User) OutputUserDetails() {
	fmt.Printf("name : %v %v \n Account Number: %v \n Balance %.2f", u.firstName, u.lastName, u.accountNumber, u.balance)

}*/
