package models

import (
	"log"
	"rate-limited-notification/database"
	"time"
)

type RateLimitCfg struct {
	Id uint	`json:"id"`
	Type string `gorm:"type:varchar(100)" json:"type"`
	MaxLimit float32 `json:"max_limit"`
	TimeInterval string `gorm:"type:varchar(100)" json:"time_interval"`
	CreationDate time.Time `json:"creation_date"`
	UpdateDate time.Time `json:"update_date"`
	Active bool `json:"active"`
}

type RateLimitCfgs []RateLimitCfg

func RateLimitCfgMigration(){
	log.Println("Rate limit migration - start")
	database.Database.AutoMigrate(&RateLimitCfg{})
	
    rateLimitsSlice := []RateLimitCfg{
        {Type: "status", MaxLimit: 2, TimeInterval: "minute", CreationDate: time.Now().UTC(), UpdateDate: time.Now().UTC(), Active: true},
        {Type: "news", MaxLimit: 1, TimeInterval: "day", CreationDate: time.Now().UTC(), UpdateDate: time.Now().UTC(), Active: true},
        {Type: "marketing", MaxLimit: 3, TimeInterval: "hour", CreationDate: time.Now().UTC(), UpdateDate: time.Now().UTC(), Active: true},
    }

    for _, rateLimit := range rateLimitsSlice {
        if err := database.Database.Create(&rateLimit).Error; err != nil {
            log.Panic("Error al insertar datos de ejemplo:", err)
        }
	}
	log.Println("Rate limit migration - end")
}