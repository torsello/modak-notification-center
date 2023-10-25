package tests

import (
	"modak-notification-center/database"
	"modak-notification-center/models"
	"testing"
)

func TestUserNotificationMigration(t *testing.T) {
	db, err := database.SetupTestDB()
	if err != nil {
		t.Fatalf("Error initializing the test database: %v", err)
	}

	database.Database = db

	models.UserNotificationMigration()

	if !db.Migrator().HasTable(&models.UserNotification{}) {
		t.Error("The UserNotification migration did not create the table correctly.")
	}
}

