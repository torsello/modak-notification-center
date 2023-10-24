package main

import (
	"log"
	"net/http"
	"rate-limited-notification/handlers"
	"rate-limited-notification/middleware"
	"rate-limited-notification/models"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	//This creates db + Initial cfg inserts
	models.RateLimitCfgMigration()
	models.UserNotificationMigration()
	log.Println("Server started at port :8080")
	mux := mux.NewRouter()
	mux.HandleFunc("/api/v1/notification", handlers.SendNotification).Methods("POST")
	//Loggin middleware
	mux.Use(middleware.LoggingMiddleware)
	//cors
	handler := cors.AllowAll().Handler(mux)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
