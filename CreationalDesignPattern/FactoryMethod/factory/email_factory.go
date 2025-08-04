package factory


import (
	"notification-factory/logger"
	"notification-factory/notifier"
)

type EmailFactory struct{}

func (EmailFactory) CreateNotifier() notifier.Notifier {
	return notifier.EmailNotifier{}
}

func (EmailFactory) CreateLogger() logger.Logger {
	return logger.EmailLogger{}
}
