
package notifier

type Notifier interface {
	Send(to, msg string)
}
