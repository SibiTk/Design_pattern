package main

import "fmt"


type Game interface {
	sports() string
}



type VolleyBall struct{}

func (v *VolleyBall)sports()string {
	return "I love VolleyBall"
}

type Cricket struct{}

func (c *Cricket)sports()string{
	return "I love Cricket"
}



func Sportstype(sports string)Game{
	switch sports {
	case "VolleyBall":
		return &VolleyBall{}
	case"Cricket":
		return &Cricket{}
		
	}
	return nil
}



func main(){
	VolleyBall :=Sportstype("VolleyBall")
	fmt.Println(VolleyBall.sports())
	Cricket :=Sportstype("Cricket")
	fmt.Println(Cricket.sports())
}


