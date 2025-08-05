package factory


import (
	"notification-factory/logger"
	"notification-factory/notifier"
)

type SMSFactory struct{}

func (SMSFactory) CreateNotifier() notifier.Notifier {
	return notifier.SMSNotifier{}
}

func (SMSFactory) CreateLogger() logger.Logger {
	return logger.SMSLogger{}
}
