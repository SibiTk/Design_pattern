package main


type State interface{
	request()
}

type TrafficSignal struct{}