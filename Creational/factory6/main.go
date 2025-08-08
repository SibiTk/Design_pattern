package main

import "fmt"

type Gateway interface{
	TransactionType() string
}

type CreditCard struct{}
func (c *CreditCard) TransactionType()string{
	return "Transaction is done By CreditCard"
}

type  Paypal struct{}
func (p *Paypal) TransactionType()string{
	return "Transaction is done by Paypal"
}




func GatewayTransaction(transactiontype string)Gateway{
switch transactiontype {
case "CreditCard":
	return &CreditCard{}
case "Paypal":
	return &Paypal{}
default:
	return nil
}
}

func main(){
CreditCard :=GatewayTransaction("CreditCard")
fmt.Println(CreditCard.TransactionType())
Paypal:=GatewayTransaction("Paypal")
fmt.Println(Paypal.TransactionType())

}