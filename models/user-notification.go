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
    Type string    `gorm:"type:varchar(100)" json:"type"`
    DeliveryDate time.Time `json:"delivery_date"`
}

type UserNotifications []UserNotification

func UserNotificationMigration(){
	log.Println("User History Notification migration")
	database.Database.AutoMigrate(&UserNotification{})
}