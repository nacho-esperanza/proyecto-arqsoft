package model

import "time"

type Booking struct {
	Id         int       `gorm:"primaryKey"`
	HotelId    int       `gorm:"foreignKey:HotelId"`
	UserId     int       `gorm:"foreignKey:UserId"`
	StartDate  time.Time `gorm:"type:date"`
	EndDate    time.Time `gorm:"type:date"`
	Price      float64   `gorm:"-"`
	TotalPrice float64   `gorm:"column:total_price"`
}

type Bookings []Booking
