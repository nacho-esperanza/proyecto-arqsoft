package app

import (
	hotelController "go-pro/controllers/hotel"
	imageController "go-pro/controllers/image"
	userController "go-pro/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {

	// Users Mapping
	router.GET("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.POST("/user", userController.UserInsert)

	// Hotels Mapping
	router.GET("/hotel/:id", hotelController.GetHotelById)
	router.GET("/hotel", hotelController.GetHotels)
	router.POST("/hotel", hotelController.HotelInsert)

	// Images Mapping
	router.POST("/image/:id", imageController.AddImage)

	log.Info("Finishing mappings configurations")
}
