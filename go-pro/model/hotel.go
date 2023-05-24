package model

type Hotel struct {
	Id     int    `gorm:"primaryKey"`
	Name   string `gorm:"type:varchar(50);not null"`
	Adress string `gorm:"type:varchar(50);not null"`
	City   string `gorm:"type:varchar(100);not null,unique"`
	Stars  int    `gorm:"type:varchar(255);not null"`
}

type Hotels []Hotel
