package main

import "fmt"


type Car interface {
    Start()
    GetName() string
}

type Bike interface {
    Start()
    GetName() string
}


type HondaCar struct{}

func (h HondaCar) Start() {
    fmt.Println("Honda car engine started")
}

func (h HondaCar) GetName() string {
    return "Honda Civic"
}

type HondaBike struct{}

func (h HondaBike) Start() {
    fmt.Println("Honda bike engine started")
}

func (h HondaBike) GetName() string {
    return "Honda CBR"
}


type ToyotaCar struct{}

func (t ToyotaCar) Start() {
    fmt.Println("Toyota car engine started")
}

func (t ToyotaCar) GetName() string {
    return "Toyota Camry"
}

type ToyotaBike struct{}

func (t ToyotaBike) Start() {
    fmt.Println("Toyota bike engine started")
}

func (t ToyotaBike) GetName() string {
    return "Toyota YBR"
}


type VehicleFactory interface {
    CreateCar() Car
    CreateBike() Bike
}


type HondaFactory struct{}

func (h HondaFactory) CreateCar() Car {
    return HondaCar{}
}

func (h HondaFactory) CreateBike() Bike {
    return HondaBike{}
}

type ToyotaFactory struct{}

func (t ToyotaFactory) CreateCar() Car {
    return ToyotaCar{}
}

func (t ToyotaFactory) CreateBike() Bike {
    return ToyotaBike{}
}

func GetVehicleFactory(brand string) VehicleFactory {
    switch brand {
    case "honda":
        return HondaFactory{}
    case "toyota":
        return ToyotaFactory{}
    default:
        return HondaFactory{}
    }
}


func main() {
    fmt.Println("=== Honda Vehicles ===")
    

    hondaFactory := GetVehicleFactory("honda")
    

    hondaCar := hondaFactory.CreateCar()
    hondaBike := hondaFactory.CreateBike()
    

    fmt.Printf("Created: %s\n", hondaCar.GetName())
    hondaCar.Start()
    
    fmt.Printf("Created: %s\n", hondaBike.GetName())
    hondaBike.Start()
    
    fmt.Println("\n=== Toyota Vehicles ===")
    

    toyotaFactory := GetVehicleFactory("toyota")
    

    toyotaCar := toyotaFactory.CreateCar()
    toyotaBike := toyotaFactory.CreateBike()
    

    fmt.Printf("Created: %s\n", toyotaCar.GetName())
    toyotaCar.Start()
    
    fmt.Printf("Created: %s\n", toyotaBike.GetName())
    toyotaBike.Start()
}
