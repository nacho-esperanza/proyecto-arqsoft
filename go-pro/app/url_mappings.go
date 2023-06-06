package app

import (
	bookingController "go-pro/controllers/booking"
	hotelController "go-pro/controllers/hotel"
	userController "go-pro/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.POST("/user", userController.UserInsert)
	router.POST("/user/login", userController.LoginUser)

	// Hotels Mapping
	router.GET("/hotel/id/:id", hotelController.GetHotelById)
	router.GET("/hotel", hotelController.GetHotels)
	router.POST("/hotel", hotelController.HotelInsert)

	// Images Mapping
	router.GET("/hotel/:hotelId/images", hotelController.GetHotelImagesById)
	router.POST("/hotel/:hotelId/image", hotelController.AddHotelImage)

	// Bookings Mapping
	router.GET("/booking/id/:id", bookingController.GetBookingById)
	router.GET("/booking", bookingController.GetBookings)
	router.POST("/booking", bookingController.CreateBooking)

	log.Info("Finishing mappings configurations")
}
