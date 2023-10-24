package models

import (
	"log"
	"rate-limited-notification/database"
	"time"
)

type UserNotification struct {
	Id uint	`json:"id"`
	User string `gorm:"type:varchar(100)" json:"user"`
    Message string `gorm:"type:varchar(100)" json:"message"`
    DeliveryDate time.Time `json:"delivery_date"`
}

type UserNotifications []UserNotification

func UserNotificationMigration(){
	log.Println("User History Notification migration - start")
	database.Database.AutoMigrate(&UserNotification{})
	log.Println("User History Notification migration - end")
}