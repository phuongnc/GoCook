package main

import "fmt"

type bankAccount struct {
	name           string
	identification string
	branch         string
	balance        int64
}

type BankAccount interface {
	Deposit(num int64)
	WithDraw(num int64)
	GetBalance() int64
}

type BankAccountBuilder interface {
	WithName(name string) BankAccountBuilder
	WithIdentification(identification string) BankAccountBuilder
	WithBranch(branch string) BankAccountBuilder
	WithBalance(balance int64) BankAccountBuilder
	Build() BankAccount
}

func init() {
	account := NewBankAccountBuilder().
		WithName("Account Holder Name").
		WithBranch("Bank Branch open account").
		WithIdentification("Identity").
		WithBalance(10).
		Build()

	account.Deposit(5)
	account.WithDraw(7)
	fmt.Println(account)
}

// implement BankAccount
func (b *bankAccount) Deposit(num int64) {
	b.balance += num
}

func (b *bankAccount) WithDraw(num int64) {
	b.balance -= num

}

func (b *bankAccount) GetBalance() int64 {
	return b.balance
}

// implement builder
func (b *bankAccount) WithName(name string) BankAccountBuilder {
	b.name = name
	return b
}
func (b *bankAccount) WithIdentification(identification string) BankAccountBuilder {
	b.identification = identification
	return b
}
func (b *bankAccount) WithBranch(branch string) BankAccountBuilder {
	b.branch = branch
	return b
}
func (b *bankAccount) WithBalance(balance int64) BankAccountBuilder {
	b.balance = balance
	return b
}
func (b *bankAccount) Build() BankAccount {
	return b
}

//init BankAccountBuilder
func NewBankAccountBuilder() BankAccountBuilder {
	return &bankAccount{}
}

// init with builder
func init() {
	account := NewBankAccountBuilder().
		WithName("Account Holder Name").
		WithBranch("Bank Branch open account").
		WithIdentification("Identity").
		WithBalance(10).
		Build()

	account.Deposit(5)
	account.WithDraw(7)
	fmt.Println(account)
}
