package logger

import "fmt"

type SMSLogger struct{}

func (SMSLogger) Log(msg string) {
	fmt.Printf("[SMS LOG]: %s\n", msg)
}

