package main

import "fmt"


type Observer interface {
    Update(message string)
}

type Subject interface {
    RegisterObserver(o Observer)
    RemoveObserver(o Observer)
    NotifyObservers()
}

type ConcreteObserver struct {
    ID int
}

func (c *ConcreteObserver) Update(message string) {
    fmt.Printf("Observer %d: Received message: %s\n", c.ID, message)
}

type ConcreteSubject struct {
    observers map[*ConcreteObserver]struct{}
    message   string
}

func (s *ConcreteSubject) RegisterObserver(o Observer) {
    obs, ok := o.(*ConcreteObserver)
    if !ok {
        return
    }
    s.observers[obs] = struct{}{}
}

func (s *ConcreteSubject) RemoveObserver(o Observer) {
    obs, ok := o.(*ConcreteObserver)
    if !ok {
        return
    }
    delete(s.observers, obs)
}

func (s *ConcreteSubject) NotifyObservers() {
    for obs := range s.observers {
        obs.Update(s.message)
    }
}


func (s *ConcreteSubject) UpdateMessage(message string) {
    s.message = message
    s.NotifyObservers()
}

func main() {
    
    subject := &ConcreteSubject{
        observers: make(map[*ConcreteObserver]struct{}),
    }


    observer1 := &ConcreteObserver{ID: 1}
    observer2 := &ConcreteObserver{ID: 2}


    subject.RegisterObserver(observer1)
    subject.RegisterObserver(observer2)

    subject.UpdateMessage("Hello, World!")

  
    subject.RemoveObserver(observer1)
    subject.UpdateMessage("Second Message")
}
