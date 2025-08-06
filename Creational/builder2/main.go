package main

import "fmt"

type Pizza struct {
	Size     string
	Toppings []string
	Cheese   bool
}

type PizzaBuilder struct {
	size     string
	toppings []string
	cheese   bool
}

func NewPizzaBuilder() *PizzaBuilder {
	return &PizzaBuilder{}
}
func (b *PizzaBuilder) SetSize(size string) *PizzaBuilder {
	b.size = size
	return b
}

func (b *PizzaBuilder) AddTopping(topping string) *PizzaBuilder {
	b.toppings = append(b.toppings, topping)
	return b
}

func (b *PizzaBuilder) WithCheese() *PizzaBuilder {
	b.cheese = true
	return b
}
func (b *PizzaBuilder) Build() *Pizza {
	return &Pizza{
		Size:     b.size,
		Toppings: b.toppings,
		Cheese:   b.cheese,
	}
}

func (p *Pizza) Describe() {
	fmt.Printf("Pizza Size: %s\n", p.Size)
	fmt.Printf("Cheese: %v\n", p.Cheese)
	fmt.Printf("Toppings: %v\n\n", p.Toppings)
}
func main() {

	pizza1 := NewPizzaBuilder().
		SetSize("Large").
		AddTopping("Chicken").
		AddTopping("Mushrooms").
		WithCheese().
		Build()

	pizza2 := NewPizzaBuilder().SetSize("Medium").AddTopping("Olives").AddTopping("Onions").Build()

	// Print pizzas
	fmt.Println("<== Pizza 1 ==>")
	pizza1.Describe()

	fmt.Println("<== Pizza 2 ==>")
	pizza2.Describe()
}