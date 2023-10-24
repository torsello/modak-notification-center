package main

import (
	"rate-limited-notification/models"
	"rate-limited-notification/services"
)

func main() {
	models.RateLimitCfgMigration()
	models.UserNotificationMigration()
	service := services.NotificationServiceImpl{Gateway: services.Gateway{}}

	service.Send("news", "user", "news 1")
	service.Send("news", "user", "news 2")
	service.Send("news", "user", "news 3")
	service.Send("news", "another user", "news 1")
	service.Send("update", "user", "update 1")
}
