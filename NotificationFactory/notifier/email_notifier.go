
package notifier

import "fmt"

type EmailNotifier struct{}

func (EmailNotifier) Send(to, msg string) {
	fmt.Printf("Email sent to %s: %s\n", to, msg)
}
