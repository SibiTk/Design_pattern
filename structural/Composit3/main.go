package main

import "fmt"


type Graphic interface {
	Draw(indent string)
}


type Circle struct {
	Name string
}

func (c *Circle) Draw(indent string) {
	fmt.Println(indent + "- Circle:", c.Name)
}


type Rectangle struct {
	Name string
}

func (r *Rectangle) Draw(indent string) {
	fmt.Println(indent + "- Rectangle:", r.Name)
}


type Group struct {
	Name     string
	Elements []Graphic
}

func (g *Group) Add(element Graphic) {
	g.Elements = append(g.Elements, element)
}

func (g *Group) Draw(indent string) {
	fmt.Println(indent + "+ Group:", g.Name)
	for _, e := range g.Elements {
		e.Draw(indent + "  ")
	}
}


func main() {

	circle1 := &Circle{Name: "Red Circle"}
	rect1 := &Rectangle{Name: "Blue Rectangle"}

	
	group1 := &Group{Name: "First Layer"}
	group1.Add(circle1)
	group1.Add(rect1)


	circle2 := &Circle{Name: "Green Circle"}
	group2 := &Group{Name: "Second Layer"}
	group2.Add(circle2)
	group2.Add(group1) 


	group2.Draw("")
}
