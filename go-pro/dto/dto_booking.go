package dto

import "time"

type BookingDto struct {
	Id         int       `json:"id"`
	HotelId    int       `json:"hotel_id"`
	UserId     int       `json:"user_id"`
	StartDate  time.Time `json:"startdate"`
	EndDate    time.Time `json:"enddate"`
	Price      float64   `json:"price"`
	TotalPrice float64   `json:"total_price"`
}

type BookingsDto struct {
	Bookings []BookingDto `json:"bookings"`
}
