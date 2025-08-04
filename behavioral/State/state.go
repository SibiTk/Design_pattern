package main

import "fmt"


type State interface {
    Request()
}


type VendingMachine struct {
    state     State
    products  int
}



type ReadyState struct {
    machine *VendingMachine
}
func (s *ReadyState) Request() {
    if s.machine.products > 0 {
        fmt.Println("Product dispensed!")
        s.machine.products--
        if s.machine.products == 0 {
            s.machine.state = &SoldOutState{s.machine}
        }
    }
}

type SoldOutState struct {
    machine *VendingMachine
}
func (s *SoldOutState) Request() {
    fmt.Println("No products left. Sold out!")
}

func main() {
    vm := &VendingMachine{products: 2}
    vm.state = &ReadyState{vm}
   
    vm.state.Request()
    vm.state.Request() 
	vm.state.Request()
}

