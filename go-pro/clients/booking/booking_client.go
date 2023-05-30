package clients

import (
	"go-pro/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func CreateBooking(booking model.Booking) model.Booking {
	result := Db.Create(&booking)

	if result.Error != nil {
		// Manejo de errores
		log.Error("Error creating booking", result.Error)
		return model.Booking{}
	}
	log.Debug("Booking Created: ", booking.Id)
	return booking
}

func GetBookingById(id int) model.Booking {
	var booking model.Booking
	Db.Where("Id = ?", id).Preload("Users").Preload("Hotels").First(&booking)
	log.Debug("Booking: ", booking)

	return booking
}

func GetBookings() []model.Booking {
	var bookings []model.Booking
	Db.Preload("Users").Preload("Hotels").Find(&bookings)
	log.Debug("Bookings: ", bookings)

	return bookings
}
