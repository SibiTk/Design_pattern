
package main

import (
	"errors"
	"fmt"
)


type Handler interface {
	SetNext(handler Handler)
	Handle(request *ATMRequest) error
}

type ATMRequest struct {
	AccountName string
	PIN         int
	Amount      int
	Balance     int 
	CorrectPIN  int
}


type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

func (b *BaseHandler) CallNext(request *ATMRequest) error {
	if b.next != nil {
		return b.next.Handle(request)
	}
	return nil
}




type AccountValidator struct {
	BaseHandler
	ValidAccount string
}

func (h *AccountValidator) Handle(request *ATMRequest) error {
	if request.AccountName != h.ValidAccount {
		return errors.New("account not recognized")
	}

	fmt.Println(" Account validated")
	return h.CallNext(request)
}


type PINValidator struct {
	BaseHandler
}

func (h *PINValidator) Handle(request *ATMRequest) error {
	if request.PIN != request.CorrectPIN {
		return errors.New("invalid PIN")
	}
	fmt.Println("pin validated")
	return h.CallNext(request)
}

// Step 3: Check balance
type BalanceChecker struct {
	BaseHandler
}

func (h *BalanceChecker) Handle(request *ATMRequest) error {
	if request.Balance < request.Amount {
		return errors.New("insufficient balance")
	}
	fmt.Println("Sufficient balance")
	return h.CallNext(request)
}


type CashDispenser struct {
	BaseHandler
}

func (h *CashDispenser) Handle(request *ATMRequest) error {
	request.Balance -= request.Amount
	fmt.Printf("Withdrawal of %d successful. New balance: %d\n", request.Amount, request.Balance)
	return h.CallNext(request)
}


func main() {
	
	request := &ATMRequest{
		AccountName: "Sibi",
		PIN:         1234,
		Amount:      500,
		Balance:     1000,
		CorrectPIN:  1234,
	}


	accountValidator := &AccountValidator{ValidAccount: "Sibi"}
	pinValidator := &PINValidator{}
	balanceChecker := &BalanceChecker{}
	cashDispenser := &CashDispenser{}


	accountValidator.SetNext(pinValidator)
	pinValidator.SetNext(balanceChecker)
	balanceChecker.SetNext(cashDispenser)

	
	err := accountValidator.Handle(request)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
