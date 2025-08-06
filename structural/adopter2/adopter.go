
package main

import (
	//"fmt"
	"github.com/sirupsen/logrus"
)

type AppLogger interface {
	Info(msg string)
	Error(msg string)
}

type LogrusLogger struct {
	logger *logrus.Logger
}

type LogrusAdapter struct {
	Adaptee *LogrusLogger
}

func (a *LogrusAdapter) Info(msg string) {
	a.Adaptee.logger.Info(msg)
}

func (a *LogrusAdapter) Error(msg string) {
	a.Adaptee.logger.Error(msg)
}


type PaymentService struct {
	logger AppLogger
}

func (p *PaymentService) ProcessPayment() {
	p.logger.Info("Payment processed successfully.")
	// simulate an error
	p.logger.Error("Payment gateway timeout.")
}

func main() {
	
	logrusLogger := &LogrusLogger{logger: logrus.New()}


	adapter := &LogrusAdapter{Adaptee: logrusLogger}

	
	service := &PaymentService{logger: adapter}
	service.ProcessPayment()
}
