package models

import (
	"log"
	"modak-notification-center/database"
	myUtils "modak-notification-center/utils"
	"time"
)

type RateLimitCfg struct {
    Id           uint      `json:"id"`
    Type         string    `gorm:"type:varchar(100)" json:"type"`
    MaxLimit     int   `json:"max_limit"`
    TimeInterval string    `gorm:"type:varchar(30)" json:"time_interval"`
    CreationDate time.Time `gorm:"type:datetime" json:"creation_date"`
    UpdateDate   time.Time `gorm:"type:datetime" json:"update_date"`
    Active       bool      `json:"active"`
}

type RateLimitCfgs []RateLimitCfg

func RateLimitCfgMigration(){
	log.Println("Rate limit migration")
	database.Database.AutoMigrate(&RateLimitCfg{})
	
	/*these inserts are made for testing, normally they should be inserted manually 
	or with some external migration, not every time you run the app, unless we look for it that way, 
	this way you could disable it with some environment variable*/

    rateLimitsSlice := []RateLimitCfg{
        {Type: "status", MaxLimit: 2, TimeInterval: myUtils.UnitsTime["minute"], CreationDate: time.Now().UTC(), UpdateDate: time.Now().UTC(), Active: true},
        {Type: "news", MaxLimit: 2, TimeInterval: myUtils.UnitsTime["day"], CreationDate: time.Now().UTC(), UpdateDate: time.Now().UTC(), Active: true},
        {Type: "marketing", MaxLimit: 3, TimeInterval: myUtils.UnitsTime["hour"], CreationDate: time.Now().UTC(), UpdateDate: time.Now().UTC(), Active: true},
    }

    for _, rateLimit := range rateLimitsSlice {
		database.Database.Model(&RateLimitCfg{}).Where("type = ? AND active = ?", rateLimit.Type, true).Update("active", false).Update("update_date", time.Now().UTC())
        if err := database.Database.Create(&rateLimit).Error; err != nil {
            log.Panic("Error inserting data:", err)
        }
	}
}