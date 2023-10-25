package tests

import (
	"modak-notification-center/database"
	"modak-notification-center/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRateLimitCfgMigration(t *testing.T) {
	db, err := database.SetupTestDB()
	if err != nil {
		t.Fatalf("Error initializing the test database: %v", err)
	}
	defer db.Migrator().DropTable(&models.RateLimitCfg{})

	database.Database = db

	models.RateLimitCfgMigration()

	if !db.Migrator().HasTable(&models.RateLimitCfg{}) {
		t.Error("The RateLimitCfg migration did not create the table correctly")
	}

	var rateLimits models.RateLimitCfgs
	db.Find(&rateLimits)

	assert.Equal(t, 3, len(rateLimits))
}

