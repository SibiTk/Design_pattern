package main

import "fmt"


type Device interface {
	On()
	Off()
	SetVolume(volume int)
}


type TV struct{}

func (t *TV) On() {
	fmt.Println("TV is ON")
}

func (t *TV) Off() {
	fmt.Println("TV is OFF")
}

func (t *TV) SetVolume(volume int) {
	fmt.Printf("TV volume set to %d\n", volume)
}

type Radio struct{}

func (r *Radio) On() {
	fmt.Println("Radio is ON")
}

func (r *Radio) Off() {
	fmt.Println("Radio is OFF")
}

func (r *Radio) SetVolume(volume int) {
	fmt.Printf("Radio volume set to %d\n", volume)
}


type RemoteControl struct {
	device Device
}

func (r *RemoteControl) TurnOn() {
	r.device.On()
}

func (r *RemoteControl) TurnOff() {
	r.device.Off()
}

func (r *RemoteControl) VolumeUp() {
	r.device.SetVolume(10)
}


type AdvancedRemote struct {
	device Device
}

func (a *AdvancedRemote) Mute() {
	a.device.SetVolume(0)
}


func main() {
	tv := &TV{}
	radio := &Radio{}

	fmt.Println("Using TV remote:")
	tvRemote := RemoteControl{device: tv}
	tvRemote.TurnOn()
	tvRemote.VolumeUp()
	tvRemote.TurnOff()

	fmt.Println("\nUsing Advanced Remote on Radio:")
	advancedRadioRemote := AdvancedRemote{device: radio}
	advancedRadioRemote.Mute()
}
