
package main

import (
	"fmt"
)


type Item struct {
	Name     string
	Quantity int
	Price    float64
}


type Iterator interface {
	HasNext() bool
	Next() *Item
}

type Cart struct {
	items []*Item
}


func (c *Cart) AddItem(item *Item) {
	c.items = append(c.items, item)
}


func (c *Cart) CreateIterator() Iterator {
	return &CartIterator{
		items: c.items,
		index: 0,
	}
}

type CartIterator struct {
	items []*Item
	index int
}

func (it *CartIterator) HasNext() bool {
	return it.index < len(it.items)
}


func (it *CartIterator) Next() *Item {
	if it.HasNext() {
		item := it.items[it.index]
		it.index++
		return item
	}
	return nil
}

func main() {

	cart := &Cart{}
	cart.AddItem(&Item{Name: "Laptop", Quantity: 1, Price: 120000})
	cart.AddItem(&Item{Name: "Headphones", Quantity: 2, Price: 1500})
	cart.AddItem(&Item{Name: "Mouse", Quantity: 1, Price: 500})

	iterator := cart.CreateIterator()

	var total float64
	for iterator.HasNext() {
		item := iterator.Next()
		fmt.Printf("Item: %s | Quantity: %d | Price: %.2f\n", item.Name, item.Quantity, item.Price)
		total += float64(item.Quantity) * item.Price
	}

	fmt.Printf("Total Cost: %.2f\n", total)
}
