package model

type account struct {
	number  string
	pwd     string
	balance float64
}

func NewAccount(number string, pwd string) *account {
	return &account{
		number: number,
		pwd:    pwd,
	}
}

func (acc *account) SetBalance(balance float64) {
	acc.balance = balance
}

func (acc *account) GetBalance() float64 {
	return acc.balance
}
