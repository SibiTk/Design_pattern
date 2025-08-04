package factory

import (
	"notification-factory/logger"
	"notification-factory/notifier"
)

type NotificationFactory interface {
	CreateNotifier() notifier.Notifier
	CreateLogger() logger.Logger
}
