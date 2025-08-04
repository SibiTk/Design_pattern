package main

import (
	"notification-factory/factory"
)

func notifyUser(factory factory.NotificationFactory, to, msg string) {
	notifier := factory.CreateNotifier()
	logger := factory.CreateLogger()

	notifier.Send(to, msg)
	logger.Log("Notification sent to " + to)
}

func main() {
	var f factory.NotificationFactory
	channel := "sms" // or "email"

	if channel == "email" {
		f = factory.EmailFactory{}
	} else {
		f = factory.SMSFactory{}
	}

	notifyUser(f, "Sibit@gmail.com", "Hello, welcome!")
}
