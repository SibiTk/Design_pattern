package main

import "fmt"

type Game interface {
	startgame()
	breakgame()
	EndGame()
	score()
}

func PlayGame(g Game) {
	g.startgame()
	g.breakgame()
	g.EndGame()
	g.score()
}

type BaseGame struct{}
func (b *BaseGame)score(){
	fmt.Println("Final Result Was Anounced")
}

type VolleyBall struct{
	BaseGame
}

func (v *VolleyBall) startgame() {
	fmt.Println("Umpire started the volley Ball game ")
}
func (v *VolleyBall) breakgame() {
	fmt.Println("Break was given for the volley Ball team")
}

func (v *VolleyBall) EndGame() {
	fmt.Println("Volley Ball tournment Was ended")
}

// func (v *VolleyBall) score() {
// 	fmt.Println("Final Result Was Anounced")
// }

type Cricket struct{
	BaseGame
}


func (c *Cricket) startgame() {
	fmt.Println("Umpire started the Cricket")
}

func (c *Cricket) breakgame() {
	fmt.Println("1st Half is Over")
}

func (c *Cricket) EndGame() {
	fmt.Println("Cricket tournament was Ended")
}

// func(c *Cricket)score(){
// 	fmt.Println("Final REsult Was Anounced")
// }

func main() {

	fmt.Println("<====VolleyBAll====>")
	PlayGame(&VolleyBall{})

	fmt.Println("<====Cricket====>")
	PlayGame(&Cricket{})
}
