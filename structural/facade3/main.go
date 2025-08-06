
package main

import "fmt"


type Account struct {
	name string
}

func NewAccount(name string) *Account {
	fmt.Println("Account created:", name)
	return &Account{name: name}
}

func (a *Account) CheckAccount(name string) bool {
	return a.name == name
}

type SecurityCode struct {
	code int
}

func NewSecurityCode(code int) *SecurityCode {
	fmt.Println("Security code set.")
	return &SecurityCode{code: code}
}

func (s *SecurityCode) CheckCode(input int) bool {
	return s.code == input
}

type Wallet struct {
	balance int
}

func NewWallet() *Wallet {
	return &Wallet{balance: 0}
}

func (w *Wallet) Credit(amount int) {
	w.balance += amount
	fmt.Printf("Credited: %d. Current Balance: %d\n", amount, w.balance)
}

func (w *Wallet) Debit(amount int) error {
	if w.balance < amount {
		return fmt.Errorf("insufficient funds")
	}
	w.balance -= amount
	fmt.Printf("Debited: %d. Current Balance: %d\n", amount, w.balance)
	return nil
}

type Notification struct{}

func (n *Notification) SendCredit(amount int) {
	fmt.Println("Notification: Amount credited:", amount)
}

func (n *Notification) SendDebit(amount int) {
	fmt.Println("Notification: Amount debited:", amount)
}

type Ledger struct{}

func (l *Ledger) RecordTransaction(txType string, amount int) {
	fmt.Printf("Ledger: %s of %d recorded\n", txType, amount)
}


type ATM struct {
	account      *Account
	securityCode *SecurityCode
	wallet       *Wallet
	notification *Notification
	ledger       *Ledger
}

func NewATM(accountName string, pin int) *ATM {
	return &ATM{
		account:      NewAccount(accountName),
		securityCode: NewSecurityCode(pin),
		wallet:       NewWallet(),
		notification: &Notification{},
		ledger:       &Ledger{},
	}
}

func (a *ATM) Withdraw(name string, pin int, amount int) error {
	fmt.Println("\nProcessing Withdrawal")
	if !a.account.CheckAccount(name) {
		return fmt.Errorf("account name mismatch")
	}
	if !a.securityCode.CheckCode(pin) {
		return fmt.Errorf("invalid PIN")
	}
	if err := a.wallet.Debit(amount); err != nil {
		return err
	}
	a.notification.SendDebit(amount)
	a.ledger.RecordTransaction("Withdrawal", amount)
	return nil
}

func (a *ATM) Deposit(name string, pin int, amount int) error {
	fmt.Println("\nProcessing Deposit")
	if !a.account.CheckAccount(name) {
		return fmt.Errorf("account name mismatch")
	}
	if !a.securityCode.CheckCode(pin) {
		return fmt.Errorf("invalid PIN")
	}
	a.wallet.Credit(amount)
	a.notification.SendCredit(amount)
	a.ledger.RecordTransaction("Deposit", amount)
	return nil
}


func main() {
	atm := NewATM("Sibi", 1234)

	err := atm.Deposit("Sibi", 1234, 1000)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = atm.Withdraw("Sibi", 1234, 300)
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = atm.Withdraw("Sibi", 1234, 800)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
