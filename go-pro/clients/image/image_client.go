package clients

import (
	"go-pro/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func AddHotelImage(image model.Image) model.Image {
	result := Db.Create(&image)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("Image Created: ", image.Id)
	return image
}

// Obtener todas las imágenes de un hotel por su ID
func GetHotelImagesById(hotelId int) model.Images {
	var images model.Images
	Db.Where("hotel_id = ?", hotelId).Find(&images)
	log.Debug("Images: ", images)
	return images
}

// Obtener las imagenes de todos los hoteles
func GetHotelImages() model.Images {
	var images model.Images
	Db.Find(&images)
	log.Debug("Images: ", images)
	return images
}
