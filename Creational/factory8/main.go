
package main

import (
	"fmt"
	"strings"
)


type Notification interface {
	Send(to string, message string)
}


type EmailNotification struct{}

func (e *EmailNotification) Send(to string, message string) {
	fmt.Printf("[EMAIL] Sent to %s: %s\n", to, message)
}


type SMSNotification struct{}

func (s *SMSNotification) Send(to string, message string) {
	fmt.Printf("[SMS] Sent to %s: %s\n", to, message)
}


type PushNotification struct{}

func (p *PushNotification) Send(to string, message string) {
	fmt.Printf("[PUSH] Sent to %s: %s\n", to, message)
}


func NotificationFactory(notificationType string) Notification {
	switch strings.ToLower(notificationType) {
	case "email":
		return &EmailNotification{}
	case "sms":
		return &SMSNotification{}
	case "push":
		return &PushNotification{}
	default:
		return nil
	}
}


type User struct {
	Name       string
	Contact    string
	NotifyType string
}

func main() {
	
	users := []User{
		
		{"Sibi", "sibit@gmail.com", "email"},
		{"Dabbu", "9789470508", "sms"},
		{"Tharshan", "tk_id", "push"},
		{"kumar", "kumar@gmail.com", "email"},
		{"Kalai", "9843091197", "sms"},
	}

	
	for i, user := range users {
		notification := NotificationFactory(user.NotifyType)
		if notification == nil {
			fmt.Printf("Unsupported notification type for %s\n", user.Name)
			continue
		}
		fmt.Print(i )
		message := fmt.Sprintf("Hello %s! This is your notification.", user.Name)
		notification.Send(user.Contact, message)
	}
}




