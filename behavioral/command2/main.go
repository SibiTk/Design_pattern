package main

import "fmt"

type Command interface{
	execute()
}

type Fan struct{}

func (f *Fan)On(){
	fmt.Println("Fan is on")
}

func (f *Fan)Off(){
	fmt.Println("Fan is Off")
}
type FanOnCommand struct{
	Fan *Fan
}

func (c *FanOnCommand) execute(){
	c.Fan.On()
}
type FanOffCommand struct{
	Fan *Fan
}

func (c *FanOffCommand)execute(){
	c.Fan.Off()
}


type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(c Command) {
	r.command = c
}

func (r *RemoteControl) PressButton() {
	r.command.execute()
}


func main(){
 fan :=&Fan{}

 FanOn :=&FanOnCommand{Fan: fan}
 FanOff :=&FanOffCommand{Fan: fan}

 remote :=&RemoteControl{}

 remote.SetCommand(FanOn)
 remote.PressButton()

  remote.SetCommand(FanOff)
 remote.PressButton()
}
