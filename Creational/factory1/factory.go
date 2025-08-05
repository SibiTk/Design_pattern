package main

import "fmt"


type Animal interface{
	speak() string
}

type Dog struct{}
func (d Dog)speak()string{
	return "woof"

}

type Cat struct{}

func (c Cat)speak()string{
	return "meow!!!!"
}
func Animalfactory(Animaltype string)Animal{
	if Animaltype =="dog"{
		return &Dog{}
	}else if  Animaltype== "cat"{
		return &Cat{}
	}

	return nil
}

func main(){
	dog:= Animalfactory("dog")
	fmt.Println(dog.speak())
	cat :=Animalfactory("cat")
	fmt.Println(cat.speak())
}