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
