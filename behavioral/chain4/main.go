package main

import "fmt"

type SupportHandler interface {
	SetNext(handler SupportHandler)
	Execute(*Employee)
}
type BaseHandler struct {
	next SupportHandler
}

type Employee struct{
	name string
	EmployeeId int
	Attendance bool
	Reason string
}
func (b *BaseHandler) SetNext(handler SupportHandler) {
	b.next = handler
}

func(b *BaseHandler)CallNext(e *Employee){
	if b.next !=nil{
		b.next.Execute(e)
	}
}


type TeamLead struct{
	BaseHandler
}

func(t *TeamLead)Execute(e *Employee){
if e.Reason=="Half Day Leave"{
	fmt.Printf("Leave is approved %s \n",e.name)
}

}
