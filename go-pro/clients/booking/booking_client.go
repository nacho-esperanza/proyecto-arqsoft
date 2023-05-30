package clients

import (
	"go-pro/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	"time"
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

/* cantidadDeReservasEnLaFechaSeleecionada =
"SELECT COUNT(*) from hotels h JOIN reservas r ON h.id = r.hotel_id
WHERE h.hotel_id = hotelID AND r.fecha_desde <= fecha_desde AND r.fecha_hasta >= fechaHasta" */
/*
func GetBookingsByDateRange(hotelId int, startDate time.Time, endDate time.Time) int {
	var bookings int
	Db.Preload("Users").Preload("Hotels").Where("HotelId = ? AND StartDate >= ? AND EndDate <= ?", hotelId, startDate, endDate).Find(&bookings)
	log.Debug("Bookings: ", bookings)

	return bookings
}*/

func GetBookingsByDateRange(hotelId int, startDate time.Time, endDate time.Time) int {
	var count int
	err := Db.Preload("Users").Preload("Hotels").Where("HotelId = ? AND StartDate >= ? AND EndDate <= ?", hotelId, startDate, endDate).Count(&count)
	if err != nil {
		// Manejo de errores
		log.Error("Error GetBookings", err)
		return int(0)
	}

	log.Debug("Number of bookings: ", count)

	return count
}

/*
func GetBookingsByDateRange(hotelId int, startDate time.Time, endDate time.Time) int {
	var count int
	Db.Model(&model.Booking{}).Where("HotelId = ? AND StartDate >= ? AND EndDate <= ?", hotelId, startDate, endDate).Count(&count)
	log.Debug("Number of bookings: ", count)

	return count
}*/
