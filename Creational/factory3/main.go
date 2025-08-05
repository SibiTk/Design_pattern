package main

import "fmt"

type logistics interface{
	transportation() string
}

type RoadTransport struct{}

func (r RoadTransport) transportation()string {
	return "Orders are Delivered by truck"
}
type SeaTransport struct{}

func (s SeaTransport) transportation()string{
	return "Orders are Delivered by Boat"
}

func logisticsFactory(logisticstype string )logistics{
	if logisticstype=="Road Transport"{
		return &RoadTransport{}
		
	}else if logisticstype=="Sea Transport"{
		return &SeaTransport{}
	}
	return nil
}
func main(){
	roadlogistics:=logisticsFactory("Road Transport")
	fmt.Println(roadlogistics.transportation())
	SeaLogistics :=logisticsFactory("Sea Transport")
	fmt.Println(SeaLogistics.transportation())
}