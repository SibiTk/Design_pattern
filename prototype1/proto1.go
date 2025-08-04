package main

import "fmt"

type Cloneable interface{
	clone() Cloneable
}

type Person struct{
	Name string
	Id int64
}

func (p *Person)clone()Cloneable{
	return &Person{
		Name: p.Name,
		Id: p.Id,
	}
}
func main(){
	person1:=&Person{Name:"sibi",Id: 101}

	clone:=person1.clone().(*Person)
	fmt.Println(person1)
	fmt.Println(clone)
	
	clone2:=clone.clone().(*Person)
	fmt.Println("Cloned Object",clone2)
}