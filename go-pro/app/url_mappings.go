package app

import (
	bookingController "go-pro/controllers/booking"
	hotelController "go-pro/controllers/hotel"
	userController "go-pro/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	log.Info("Starting mappings configurations")

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)

	// Only for Admin
	router.GET("/user", userController.GetUsers)

	router.POST("/user", userController.UserInsert)
	router.POST("/user/login", userController.LoginUser)

	// Hotels Mapping
	router.GET("/hotel/id/:id", hotelController.GetHotelById)
	router.GET("/hotel", hotelController.GetHotels)

	// Only for Admin
	router.POST("/hotel", hotelController.HotelInsert)

	// Images Mapping
	router.GET("/hotel/:hotelId/images", hotelController.GetHotelImagesById)
	router.GET("/hotel/images", hotelController.GetHotelImages)
	router.POST("/hotel/:hotelId/image", hotelController.AddHotelImage)

	// Bookings Mapping

	// Only for Admin
	router.GET("/booking/id/:id", bookingController.GetBookingById)

	// Only for Admin
	router.GET("/booking", bookingController.GetBookings)

	router.GET("/booking/user/:userId", bookingController.GetBookingsByUserId)
	router.POST("/booking", bookingController.CreateBooking)

	log.Info("Finishing mappings configurations")
}
