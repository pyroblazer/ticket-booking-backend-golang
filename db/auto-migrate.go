package db

import (
	"github.com/pyroblazer/ticket-booking-backend-golang/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{}, &models.Ticket{}, &models.User{})
}
