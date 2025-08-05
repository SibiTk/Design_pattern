package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
	name string
}

var (
	instance *Singleton
	instance1 *Singleton
	once     sync.Once
	//lock = &sync.Mutex{}
)

func GetInstance() *Singleton {
	 once.Do(func() {

		instance = &Singleton{name: "MySingletonInstance"}
		fmt.Println("Singleton instance created.")
	})
	return instance
}

//Example method for the Singleton
func (s *Singleton) GetName() string {
	return s.name
}

// func GetInstance1() *Singleton {

// 	instance = &Singleton{name: "My name is Sibi"}
// 	fmt.Println("Singleton Instance is Created again")
// 	return instance
// }


func GetInstance1() *Singleton {
   
        if instance1 == nil {
			instance1 = &Singleton{name: "My name is Sibi"}
        } else {
            fmt.Println("Single instance already created.")
        }
    return instance1
}



func (s *Singleton) PrintAddress() {
	fmt.Println("Address ==> ", &s)
}



func (s *Singleton) GetPlace() string {
	return s.name
}
