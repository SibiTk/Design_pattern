package main

import "fmt"


type Pricing interface {
	GetPrice() int
}


type Product struct {
	Name  string
	Price int
}

func (p Product) GetPrice() int {
	return p.Price
}


type ProductBinding struct {
	Items    []Pricing
	Discount int
}

func (pb *ProductBinding) Add(item Pricing) {
	pb.Items = append(pb.Items, item)
}

func (pb ProductBinding) GetPrice() int {
	total := 0
	for _, item := range pb.Items {
		total += item.GetPrice()
	}
	return total - pb.Discount
}


type Order struct {
	DeliveryPrice int
	Items         []Pricing
}

func (o *Order) Add(item Pricing) {
	o.Items = append(o.Items, item)
}

func (o Order) GetPrice() int {
	total := o.DeliveryPrice
	for _, item := range o.Items {
		total += item.GetPrice()
	}
	return total
}

func main() {

	phone := Product{Name: "iPhone", Price: 35590}
	charger := Product{Name: "iPhone Charger", Price: 12200}
	airpods := Product{Name: "AirPods", Price: 15900}

	bundle := &ProductBinding{Discount: 2000}
	bundle.Add(charger)
	bundle.Add(airpods)

	order := &Order{DeliveryPrice: 1500}
	order.Add(phone)
	order.Add(bundle)

	fmt.Printf("Total price: %d\n", order.GetPrice()) 
}

