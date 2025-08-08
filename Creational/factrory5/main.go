package main 

import "fmt"

type Vehicle interface{
	drive() string
}

type Car struct{}

func (c *Car)drive() string{
	//fmt.Println("Driving Car")
	 return "Driving Car"
}


type Bike struct{}

func (b *Bike)drive()string{
	//fmt.Println("Driving Bike")
	return "Driving Bike"
}

type Truck struct{}
func (t *Truck)drive()string{
	return "Driving"
}

func VehicleFactory(vehicletype string)Vehicle{
switch vehicletype {
	case "car":
		return &Car{}
	case "bike":
		return &Bike{}
	case "truck":
		return &Truck{}
	default:
		return nil
	}
}

func main(){
	car :=VehicleFactory("car")
	fmt.Println(car.drive())

	Bike:=VehicleFactory("Bike")
	fmt.Println(Bike.drive())

}