

package main

import "fmt"


type SupportHandler interface {
	SetNext(handler SupportHandler)
	HandleRequest(issue string)
}

type BaseHandler struct {
	next SupportHandler
}

func (b *BaseHandler) SetNext(handler SupportHandler) {
	b.next = handler
}

func (b *BaseHandler) CallNext(issue string) {
	if b.next != nil {
		b.next.HandleRequest(issue)
	} else {
		fmt.Println("Issue could not be resolved.")
	}
}


type Level3Support struct {
	BaseHandler
}

func (l1 *Level3Support) HandleRequest(issue string) {
	if issue == "password reset" {
		fmt.Println("L3 Hnadled: password reset")
	} else {
		fmt.Println("L3 Handled: Can't handle, escalating...")
		l1.CallNext(issue)
	}
}


type Level2Support struct {
	BaseHandler
}

func (l2 *Level2Support) HandleRequest(issue string) {
	if issue == "account locked" {
		fmt.Println("l2 handled: account locked")
	} else {
		fmt.Println("L2 : Can't handle, escalating.. .")
		l2.CallNext(issue)
	}
}


type Level1Support struct {
	BaseHandler
}

func (m *Level1Support) HandleRequest(issue string) {
	if issue == "refund request" {
		fmt.Println("L1 handled: refund request")
	} else {
		fmt.Println("L1: Can't handle this issue.")
		m.CallNext(issue)
	}
}


func main() {
	
	L3 := &Level3Support{}
	L2 := &Level2Support{}
	L1 := &Level1Support{}


	L3.SetNext(L2)
	L2.SetNext(L1)


	fmt.Println("Customer Issue: password reset")
	L3.HandleRequest("password reset")

	fmt.Println("\nCustomer Issue: account locked")
	L3.HandleRequest("account locked")

	fmt.Println("\nCustomer Issue: refund request")
	L3.HandleRequest("refund request")

	fmt.Println("\nCustomer Issue: unknown issue")
	L3.HandleRequest("technical failure")
}
