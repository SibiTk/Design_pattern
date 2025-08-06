package main


import "fmt"


type Colour interface {
	PrintColour()
}


type Red struct{}
func (r *Red) PrintColour()   { fmt.Println("I'm Red!") }

type Blue struct{}
func (b *Blue) PrintColour()  { fmt.Println("I'm Blue!") }


type Shape interface {
	GetColour()
	SetColour(Colour)
}

type Square struct {
	colour Colour
}

func (s *Square) GetColour() {
	fmt.Print("Square says: ")
	s.colour.PrintColour()
}
func (s *Square) SetColour(c Colour) {
	s.colour = c
}

type Circle struct {
	colour Colour
}

func (c *Circle) GetColour() {
	fmt.Print("Circle says: ")
	c.colour.PrintColour()
}
func (c *Circle) SetColour(co Colour) {
	c.colour = co
}

func main() {
	r := &Red{}
	b := &Blue{}

	s := &Square{}
	s.SetColour(r)
	s.GetColour()
	s.SetColour(b)
	s.GetColour() 

	cir := &Circle{}
	cir.SetColour(b)
	cir.GetColour() 
	cir.SetColour(r)
	cir.GetColour() 
}
