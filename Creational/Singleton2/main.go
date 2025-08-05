package main

import "fmt"

type Singleton struct{
	value int
}

var instance *Singleton
func GetInstance() *Singleton{
	if instance ==nil{
		instance=&Singleton{value:10}
	}
	return instance
}
func main(){
	s1 := GetInstance()
	s2 := GetInstance()

	fmt.Println("s1 value =", s1.value)
	fmt.Println("s2 value =", s2.value)
	fmt.Println("same instance?", s1 == s2)

}