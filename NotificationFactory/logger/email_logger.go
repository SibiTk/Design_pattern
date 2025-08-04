package logger


import "fmt"

type EmailLogger struct{}

func (EmailLogger) Log(msg string) {
	fmt.Printf("[EMAIL LOG]: %s\n", msg)
}
