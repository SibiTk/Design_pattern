package main

import (
 "fmt"
)


type Cloneable interface {
  Clone() Cloneable
 }


type Product struct {
  name     string
  category string
 }


func (p *Product) Clone() Cloneable {
 return &Product{name: p.name, category: p.category}
}

func (p *Product) SetName(name string) {
 p.name = name
}

func (p *Product) GetDetails() string {
 return fmt.Sprintf("Product Name: %s, Category: %s", p.name, p.category)
}

func main() {

 original := &Product{name: "1+", category: "Electronics"}
 fmt.Println(original.GetDetails()) 


 cloned := original.Clone().(*Product)
 cloned.SetName("Smartphocne")
 fmt.Println(cloned.GetDetails()) 

}