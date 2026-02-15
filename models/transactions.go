package models

func (u *User) Deposit(amount float64) (*User, error) {

	u.Balance += amount
	u.SaveAccount()
	return u, nil
}
func (u *User) Withdrawal(amount float64) (*User, error) {

	u.Balance -= amount
	u.SaveAccount()
	return u, nil
}
