package clients

import (
	"go-pro/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetHotelById(id int) model.Hotel {
	var hotel model.Hotel

	Db.Where("Id = ?", id).Preload("Images").First(&hotel)
	log.Debug("Hotel: ", hotel)

	return hotel
}

func GetHotels() model.Hotels {
	var hotels model.Hotels

	// Recordar poner los preload para que traiga los datos de las tablas relacionadas

	Db.Preload("Images").Find(&hotels)

	Db.Find(&hotels)

	log.Debug("Hotels: ", hotels)

	return hotels
}

func InsertHotel(hotel model.Hotel) model.Hotel {
	result := Db.Create(&hotel)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("Error creating hotel:", result.Error)

		return model.Hotel{}
	}
	log.Debug("Hotel Created: ", hotel.Id)
	return hotel
}

func AddHotelImage(image model.Image) model.Image {
	result := Db.Create(&image)
	if result.Error != nil {
		// Manejo de errores
		log.Error("Error creating hotel image:", result.Error)
		return model.Image{}
	}
	log.Debug("Hotel Image Created: ", image.Id)
	return image
}
