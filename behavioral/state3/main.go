package main

import "fmt"


type State interface {
	Play()
	Pause()
	Stop()
}


type MediaPlayer struct {
	state State
}

func (m *MediaPlayer) SetState(state State) {
	m.state = state
}

func (m *MediaPlayer) Play() {
	m.state.Play()
}

func (m *MediaPlayer) Pause() {
	m.state.Pause()
}

func (m *MediaPlayer) Stop() {
	m.state.Stop()
}


type PlayingState struct {
	player *MediaPlayer
}

func (s *PlayingState) Play() {
	fmt.Println("Already playing.")
}

func (s *PlayingState) Pause() {
	fmt.Println("Pausing playback.....")
	s.player.SetState(&PausedState{player: s.player})
}

func (s *PlayingState) Stop() {
	fmt.Println("Stopping playback...")
	s.player.SetState(&StoppedState{player: s.player})
}


type PausedState struct {
	player *MediaPlayer
}


func (s *PausedState) Play() {
	fmt.Println("Resuming playback...")
	s.player.SetState(&PlayingState{player: s.player})
}


func (s *PausedState) Pause() {
	fmt.Println("Already paused.")
}

func (s *PausedState) Stop() {
	fmt.Println("Stopping from pause...")
	s.player.SetState(&StoppedState{player: s.player})
}


type StoppedState struct {
	player *MediaPlayer
}

func (s *StoppedState) Play() {
	fmt.Println("Starting playback......")
	s.player.SetState(&PlayingState{player: s.player})
}

func (s *StoppedState) Pause() {
	fmt.Println("Cannot pause. Media is stopped.")
}

func (s *StoppedState) Stop() {
	fmt.Println("Already stopped.")
}


func main() {
	player := &MediaPlayer{}
	stopped := &StoppedState{player: player}
	player.SetState(stopped)
	player.Play()   
	player.Pause()  
	player.Play()  
	player.Stop()  
	player.Pause()  
	player.Stop()   
}
