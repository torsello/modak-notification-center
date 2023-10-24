package database

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database = func () (db *gorm.DB){
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic(errEnv) 
	}
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	
	dsn := username+":"+password+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8mb4"
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		panic(err) 
	}else{
		return db
	}
}()