package model

type Image struct {
	Id      int    `gorm:"primaryKey"`
	Url     string `gorm:"type:varchar(255);not null"`
	HotelId int    `gorm:"foreignKey:HotelId"`
}

type Images []Image
