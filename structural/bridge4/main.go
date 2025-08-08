package main

import "fmt"


type MessageSender interface {
	SendMessage(message, to string)
}


type EmailSender struct{}

func (e *EmailSender) SendMessage(message, to string) {
	fmt.Printf("Sending Email to %s: %s\n", to, message)
}

type SMSSender struct{}

func (s *SMSSender) SendMessage(message, to string) {
	fmt.Printf("Sending SMS to %s: %s\n", to, message)
}


type Notification struct {
	sender MessageSender
}

func (n *Notification) Notify(message, to string) {
	n.sender.SendMessage(message, to)
}

type UrgentNotification struct {
	sender MessageSender
}

func (u *UrgentNotification) Notify(message, to string) {
	urgentMsg := "[URGENT] " + message
	u.sender.SendMessage(urgentMsg, to)
}


func main() {
	email := &EmailSender{}
	sms := &SMSSender{}

	normal := &Notification{sender: email}
	urgent := &UrgentNotification{sender: sms}

	normal.Notify("Meeting at 10am", "tksibi.com")
	urgent.Notify("Server down!", "123456789")
}



