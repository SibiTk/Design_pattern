
package main

import "fmt"


type Coffee interface {
	GetCost() float64
	GetDescription() string
}


type BasicCoffee struct{}
func (c *BasicCoffee) GetCost() float64{
	return 5
}

func (c *BasicCoffee) GetDescription() string {
	return "Base Coffee"
}

// Decorator struct
type CoffeeDecorator struct {
	coffee Coffee
}

func (d *CoffeeDecorator) GetCost() float64 {
	return d.coffee.GetCost()
}

func (d *CoffeeDecorator) GetDescription() string {
	return d.coffee.GetDescription()
}

// Milk decorator
type Milk struct {
	CoffeeDecorator
}

func NewMilk(coffee Coffee) *Milk {
	return &Milk{CoffeeDecorator{coffee}}
}

func (m *Milk) GetCost() float64 {
	return m.coffee.GetCost() + 5
}

func (m *Milk) GetDescription() string {
	return m.coffee.GetDescription() + ", Milk"
}

// Sugar decorator
type Sugar struct {
	CoffeeDecorator
}

func NewSugar(coffee Coffee) *Sugar {
	return &Sugar{CoffeeDecorator{coffee}}
}

func (s *Sugar) GetCost() float64 {
	return s.coffee.GetCost() + 5
}

func (s *Sugar) GetDescription() string {
	return s.coffee.GetDescription() + ", Sugar"
}

func main() {
	var myCoffee Coffee = &BasicCoffee{}

	fmt.Println(myCoffee.GetDescription(), "cost:", myCoffee.GetCost())


	myCoffee = NewMilk(myCoffee)
	fmt.Println(myCoffee.GetDescription(), "cost:", myCoffee.GetCost())

     myCoffee =NewSugar(myCoffee)
	 fmt.Println(myCoffee.GetDescription(),"Coffee:",myCoffee.GetCost())
}
