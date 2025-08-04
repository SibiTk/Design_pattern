package notifier


import "fmt"

type SMSNotifier struct{}

func (SMSNotifier) Send(to, msg string) {
	fmt.Printf("SMS sent to %s: %s\n", to, msg)
}
