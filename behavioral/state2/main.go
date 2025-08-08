package main

import "fmt"


type State interface {
	Handle(light *TrafficLight)
}

type TrafficLight struct {
	state State
}

func (t *TrafficLight) SetState(state State) {
	t.state = state
}


func (t *TrafficLight) Request() {
	t.state.Handle(t)
}


type RedLightState struct{}

func (r *RedLightState) Handle(light *TrafficLight) {
	fmt.Println("red Light - STOP")
	light.SetState(&GreenLightState{}) 
}


type GreenLightState struct{}

func (g *GreenLightState) Handle(light *TrafficLight) {
	fmt.Println("green Light - GO")
	light.SetState(&YellowLightState{}) 
}


type YellowLightState struct{}

func (y *YellowLightState) Handle(light *TrafficLight) {
	fmt.Println("yellow Light - CAUTION")
	light.SetState(&RedLightState{}) 
}


func main() {
	light := &TrafficLight{state: &RedLightState{}} 
	for i := 0; i < 6; i++ {
		light.Request()
	}
}


