package main

import "fmt"

type House struct{
Name string
Phone string
Balconi string
Floor int
Garden string 


}

type HouseBuilder struct{
name string
phone string
balconi  string
floor int
garden string
House

}

func NewHouseBuilder(name,phone string)*HouseBuilder{
	return &HouseBuilder{
		name :name,
		phone :phone,
		
	}
}

func(h *HouseBuilder)SetBalconi(Balconi string)*HouseBuilder{
	h.balconi=Balconi
	return h
}

func (h *HouseBuilder)SetFloor(floor int)*HouseBuilder{
h.floor =floor
return h
}

func (h *HouseBuilder)SetGarden(garden string)*HouseBuilder{
h.garden =garden
return h
}

func (h *HouseBuilder)Build()*House{
	return &House{
		Name: h.name,
		Phone: h.phone,
		Balconi: h.balconi,
		Floor: h.floor,
		Garden: h.garden,
	}
}


func (c *House)Print(){
	fmt.Println("<== DETAILS OF OWNER AND HOUSE=cd ..==>")
	fmt.Printf("Name : %s\n",c.Name)
	fmt.Printf("Phone : %s\n",c.Phone)
	fmt.Printf("Balconi :%s\n",c.Balconi)
	fmt.Printf("Floor :%d\n",c.Floor)

}


func main(){

owner1 :=NewHouseBuilder("Dharmesh","9344954272").SetBalconi("Found in the first Floor").SetFloor(2).SetGarden("Garden is on the right side of the house").Build()
owner1.Print()

}