package main

import "fmt"

type ChefReceivedOrder struct{}

func(c *ChefReceivedOrder) kitchen(){
	fmt.Println("Chef taken the order and Get the vegetables and  meat to cook the food")
}

type Preperation struct{}

func(c *Preperation) Processing(){
	fmt.Println("Started to cook the Food")
}

type FoodDelivered struct{}

func(c *FoodDelivered)CustomerReceived(){
	fmt.Println("Customer Received the food and started to eat")
}

type Restaurant struct{
	chefReceivedOrder *ChefReceivedOrder
	preperation    *Preperation
	foodDelivered  *FoodDelivered
}
func NewRestaurant()*Restaurant{
	return &Restaurant{
    chefReceivedOrder: &ChefReceivedOrder{},
	preperation :   &Preperation{},
	foodDelivered : &FoodDelivered{},
	}
}

	func (r *Restaurant)CustomerOrder() {
		r.chefReceivedOrder.kitchen()
		r.preperation.Processing()
		r.foodDelivered.CustomerReceived()
	
	}



	func main(){
		restaurant :=NewRestaurant()
		fmt.Println("Eating Delicious Food in Restaurant")
		restaurant.CustomerOrder()
	}


