package main

import "fmt"

type PaymentProcessor interface {
	Validate()
	ConnectGateway()
	ProcessPayment()
	
}

func MakePayment(p PaymentProcessor) {
	p.Validate()
	p.ConnectGateway()
	p.ProcessPayment()

}

type CreditCard struct{}

func (c *CreditCard)Validate(){
	fmt.Println("Validating Credit card here")
}
func (c *CreditCard)ConnectGateway(){
	fmt.Println("Connecting gate way to process the transaction in Cred card")
}

func(c *CreditCard)ProcessPayment(){
	fmt.Println("Your Credit card payment is Succesful")
}

type UPI struct{}
func (u *UPI)Validate(){
	fmt.Println("Validate the Upi password")
}

func(u *UPI) ConnectGateway(){
	fmt.Println("Connect the gateway to process the upi transaction")
}

func(u *UPI)ProcessPayment(){
	fmt.Println("Your Upi Payment is successful")
}

type Paypal struct{}



func (p *Paypal)Validate(){
	fmt.Println("Validate the Paypal password")
}

func(p *Paypal) ConnectGateway(){
	fmt.Println("Connect the gateway to process the PAYPAL transaction")
}

func(p *Paypal)ProcessPayment(){
	fmt.Println("Your Paypal Payment is successful")
}

func main(){
	fmt.Println("<===Credit Card===>")
	MakePayment(&CreditCard{})
	fmt.Println("<===UPI===>")
	MakePayment((&UPI{}))
	fmt.Println("<==PayPal===>")
	MakePayment((&Paypal{}))

}