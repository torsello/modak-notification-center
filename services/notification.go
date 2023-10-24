package services

type NotificationService interface {
	Send(notificationType, user, message string)
}

type NotificationServiceImpl struct {
	Gateway Gateway
}

func (service NotificationServiceImpl) Send(typeName, user, message string) {
	service.Gateway.Send(user, message)
}