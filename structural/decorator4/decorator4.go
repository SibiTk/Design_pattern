package main
import "fmt"

type Car interface {
	GetCost() int
	GetDescription() string
}

type NewCar struct{}

func (c *NewCar) GetCost() int {
	return 200000000
}
func (c *NewCar) GetDescription() string {
	return "Buying Porche 911"
}

type CarDecorator struct {
	car Car
}

func (s *CarDecorator) GetCost() int {
	return s.car.GetCost()
}

func (s *CarDecorator) GetDescription() string {
	return s.car.GetDescription()
}


type SafetyClothes struct{
	CarDecorator
}

func NewSafetyClothes(car Car) *SafetyClothes{
   return &SafetyClothes{}
}

func (s *SafetyClothes)GetCost()int{
	return s.car.GetCost()+5000
}


func (s *SafetyClothes)GetDescription()string{
	return s.car.GetDescription()+" Safety clothes are added"
}


func main(){

	var MyCar Car =&NewCar{}
	fmt.Println(MyCar.GetDescription(),"Cost :",MyCar.GetCost())
	MyCar =NewSafetyClothes(MyCar)
	fmt.Println(MyCar.GetDescription(),"Total Cost :",MyCar.GetCost())
}