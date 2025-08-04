package main

import (
	"fmt"
)

func main() {
	// Get the Singleton instance multiple times
	instance1 := GetInstance()

	instance2 := GetInstance()
	instance3 := GetInstance()
	 instance4 := GetInstance1()

	

	fmt.Println("Instance 1 Name:", instance1.GetName())
	fmt.Println("Instance 1.1 Name:", instance1.GetName())

	fmt.Println("Instance 2 Name:", instance2.GetName())
	fmt.Println("Instance 3 Name:", instance3.GetName())
	fmt.Println("Instance 4 Name:", instance4.GetPlace())

	if instance1 == instance2 && instance2 == instance3 {
		fmt.Println("All instances are the same Singleton works!")
	} else {
		fmt.Println("Singleton pattern is broken.")
	}
	if instance4 != instance3 {
		fmt.Println("Instance is Created Again")
	}

	instanceX := GetInstance1()
	instanceY := GetInstance1()

	instanceX.PrintAddress()
	instanceY.PrintAddress()

	instanceA := GetInstance()
	instanceB := GetInstance()

	instanceA.PrintAddress()
	instanceB.PrintAddress()

	instanceX.name = "instanceX"
	instanceY.name = "instanceY"
	instanceA.name = "instanceA"
	instanceB.name = "instanceB"

	

	fmt.Println(instanceX.name)
	fmt.Println(instanceY.name)
	fmt.Println(instanceA.name)
	fmt.Println(instanceB.name)

}
