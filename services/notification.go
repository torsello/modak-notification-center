package services

import (
	"log"
	"modak-notification-center/database"
	"modak-notification-center/models"
	myUtils "modak-notification-center/utils"
	"strings"
	"sync"
	"time"
)

type NotificationService interface {
	Send(notificationType, user, message string)
}

type NotificationServiceImpl struct {
	Gateway Gateway
}

func (service NotificationServiceImpl) Send(typeName, user, message string) bool {
	sentOk := false
	log.Println("User Notification - Validation - User: "+user+" Type: "+typeName)
	wg := &sync.WaitGroup{}

	rateLimitChan := make(chan models.RateLimitCfg)
	
	userNotificationsChan := make(chan models.UserNotifications)
	errorUserNotificationChan := make(chan error)

	wg.Add(2)
	go getRateLimitCfg(typeName, rateLimitChan, wg)
	rateLimitCfg := <- rateLimitChan
	go checkForNotifications(typeName, user, rateLimitCfg, userNotificationsChan, errorUserNotificationChan, wg)
	wg.Wait()
	
	userNotifications := <- userNotificationsChan
	if len(userNotifications) < rateLimitCfg.MaxLimit || rateLimitCfg.Id <= 0 {
		saveNotification(typeName, user, message)
		service.Gateway.Send(user, message)
		sentOk = true
	}
	
	return sentOk
}

func getRateLimitCfg(typeName string, resultChan chan<- models.RateLimitCfg, wg *sync.WaitGroup) {
	rateLimitCfg := models.RateLimitCfg{}
    if database.Database.Where("type = ?", typeName).Where("active = ?", true).Limit(1).Find(&rateLimitCfg).RowsAffected == 0 {
		log.Println("CFG for "+ typeName + " not exists, we proceed to send the notification without limit")
	}
	wg.Done()
    resultChan <- rateLimitCfg
}

func checkForNotifications(typeName string, user string, rateLimitCfg models.RateLimitCfg, resultChan chan<- models.UserNotifications, errorChan chan<- error, wg *sync.WaitGroup) {
	userNotifications := models.UserNotifications{}

	if rateLimitCfg.Id > 0 {
		validationDate := myUtils.GetValidationDate(time.Now().UTC(), strings.ToLower(rateLimitCfg.TimeInterval))
		if err := database.Database.
		Where("type = ? AND user = ? AND delivery_date > ?", typeName, user, validationDate).
		Order("delivery_date DESC").
		Limit(rateLimitCfg.MaxLimit).
		Find(&userNotifications).
        Error; err != nil {
			wg.Done()
			errorChan <- err
			close(resultChan)
			return
    	}
	}

	wg.Done()
    resultChan <- userNotifications
}

func saveNotification(typeName, user, message string){
	notification := models.UserNotification{
		User: user,
		Message: message, 
		Type: typeName,
		DeliveryDate: time.Now().UTC(),
	}

	database.Database.Save(&notification)
}