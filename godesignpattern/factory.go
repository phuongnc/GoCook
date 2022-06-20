package main

import "fmt"

type PaymentMethod interface {
	Pay(amount int64)
	SetBalance(balance int64)
}

type CreditAccount struct {
	balance int64
}
type DebitAccount struct {
	balance int64
}

func (c *CreditAccount) Pay(amount int64) {
	c.balance -= amount

}
func (c *CreditAccount) SetBalance(balance int64) {
	c.balance = balance
}

func (d *DebitAccount) Pay(amount int64) {
	d.balance -= (amount + 1)

}
func (d *DebitAccount) SetBalance(balance int64) {
	d.balance = balance
}

//Factory using PaymentMethod to init CreditAccount or DebitAccount base on Type
func GetPaymentMethod(t string) PaymentMethod {
	switch t {
	case "credit":
		return new(CreditAccount)
	default:
		return new(DebitAccount)
	}
}

func init() {
	credit := GetPaymentMethod("credit")
	credit.SetBalance(100)
	debit := GetPaymentMethod("debit")
	debit.SetBalance(100)

	credit.Pay(10)
	debit.Pay(10)

	fmt.Println(credit)
	fmt.Println(debit)
}
