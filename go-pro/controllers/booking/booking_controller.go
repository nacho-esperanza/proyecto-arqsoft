package bookingController

import (
	"go-pro/dto"
	service "go-pro/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetBookingById(c *gin.Context) {
	log.Debug("Booking id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var bookingDto dto.BookingDto

	bookingDto, err := service.BookingService.GetBookingById(id) // se utiliza el userDto referenciandose del id

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, bookingDto)
}

func GetBookings(c *gin.Context) {
	var bookingsDto dto.BookingsDto
	bookingsDto, err := service.BookingService.GetBookings()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, bookingsDto)
}

func GetBookingsByUserId(c *gin.Context) {

	userId, _ := strconv.Atoi(c.Param("userId"))
	log.Debug("User id to load: " + c.Param("userId"))
	var bookingsDto dto.BookingsDto
	bookingsDto, err := service.BookingService.GetBookingsByUserId(userId)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, bookingsDto)
}

func CreateBooking(c *gin.Context) {
	var bookingDto dto.BookingDto
	err := c.BindJSON(&bookingDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	bookingDto, er := service.BookingService.CreateBooking(bookingDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, bookingDto)
}
