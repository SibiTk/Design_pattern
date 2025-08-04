package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64)
}

type CreditCard struct {
	Name   string
	Number string
	CVV    string
}

func (c *CreditCard) Pay(amount float64) {
	fmt.Printf("Paid %.2f using Credit Card (%s).\n", amount, c.Name)
}


type PayPal struct {
	Email string
}

func (p *PayPal) Pay(amount float64) {
	fmt.Printf("Paid %.2f using PayPal (%s).\n", amount, p.Email)
}


type ShoppingCart struct {
	paymentMethod PaymentStrategy
}

func (c *ShoppingCart) SetPaymentMethod(p PaymentStrategy) {
	c.paymentMethod = p
}

func (c *ShoppingCart) Checkout(amount float64) {
	if c.paymentMethod == nil {
		fmt.Println("Payment method not set!")
		return
	}
	c.paymentMethod.Pay(amount)
}

func main() {
	cart := &ShoppingCart{}


	creditCard := &CreditCard{Name: "Sibi", Number: "1234-5678-9012", CVV: "123"}
	cart.SetPaymentMethod(creditCard)
	cart.Checkout(100.50)

	
	paypal := &PayPal{Email: "tksibi@gmail.com"}
	cart.SetPaymentMethod(paypal)
	cart.Checkout(250.00)
}
