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
/*
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
}*/

// CONTAR DIA POR DIA LAS RESERVAS PARA EL FINAL
// CAMBIAR LA QUERY
/*
func GetBookingsByDateRange(hotelId int, startDate time.Time, endDate time.Time) int {
	var count int

	txn := Db.Model(&model.Bookings{}).Where("hotel_id = ? AND start_date >= ? AND end_date <= ?", hotelId, startDate, endDate).Count(&count)
	if txn.Error != nil {
		// Manejo de errores
		log.Error("Error GetBookings", txn.Error)
		return 100000
	}

	log.Debug("Number of bookings: ", count)

	return count
}*/

// PRUEBA DE CONTAR DIA POR DIA LAS RESERVAS
// POR AHORA FUNCIONA LA QUERY NUEVA

func GetBookingsByDateRange(hotelId int, startDate time.Time, endDate time.Time) int {
	var count int

	// Obtener la lista de fechas dentro del rango
	dates := getDatesInRange(startDate, endDate)

	// Contar la cantidad de habitaciones reservadas por cada fecha
	for _, date := range dates {
		txn := Db.Model(&model.Bookings{}).Where("hotel_id = ? AND start_date <= ? AND end_date >= ?", hotelId, date, date).Count(&count)
		if txn.Error != nil {
			// Manejo de errores
			log.Error("Error GetBookings", txn.Error)
			return 100000
		}
	}

	log.Debug("Number of bookings: ", count)

	return count
}

// Función auxiliar para obtener la lista de fechas dentro de un rango
func getDatesInRange(startDate time.Time, endDate time.Time) []time.Time {
	var dates []time.Time
	currentDate := startDate

	for !currentDate.After(endDate) {
		dates = append(dates, currentDate)
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return dates
}

func GetBookingsByUserId(userId int) []model.Booking {
	var bookings []model.Booking
	// Db.Raw("SELECT * FROM bookings WHERE user_id = ?", userId).Scan(&bookings)  // esto es lo mismo que la linea de abajo PERO EN SQL
	Db.Preload("Users").Preload("Hotels").Where("user_id = ?", userId).Find(&bookings)
	log.Debug("Bookings: ", bookings)

	return bookings
}
