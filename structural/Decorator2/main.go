package main

import "fmt"


type Phone interface{
	PhoneCost() int
	PhoneModel() string
}

type OnlyPhone struct{}
func (p *OnlyPhone)PhoneCost()int{
	return 30000
}
func (p *OnlyPhone)PhoneModel()string{
	return "samsung Galaxy m31"
}

type PhoneDecorator struct{
	phone Phone
}

func(d *PhoneDecorator)PhoneCost()int{
	return d.phone.PhoneCost()
}

func(d *PhoneDecorator)PhoneModel()string{
	return d.phone.PhoneModel()
}

type Template struct{
	PhoneDecorator
}

func NewTemplate(phone Phone) *Template{
	return &Template{PhoneDecorator{phone}}
}

func (t *Template)PhoneCost()int{
return t.phone.PhoneCost()+200	
}

func (t *Template)PhoneModel()string{
	return t.phone.PhoneModel()+", Template is Added "
}


func main(){
	var myphone Phone =&OnlyPhone{}
	fmt.Println(myphone.PhoneModel(),"cost: ",myphone.PhoneCost())

	myphone =NewTemplate(myphone)
	fmt.Println(myphone.PhoneModel(),"cost:",myphone.PhoneCost())
}