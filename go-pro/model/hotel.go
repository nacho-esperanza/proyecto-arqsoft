package model

type Hotel struct {
	Id          int     `gorm:"primaryKey"`
	Name        string  `gorm:"type:varchar(50);not null"`
	Adress      string  `gorm:"type:varchar(100);not null"`
	City        string  `gorm:"type:varchar(100);not null,unique"`
	Stars       int     `gorm:"type:int;not null"`
	Description string  `gorm:"type:varchar(255);not null"`
	Price       float32 `gorm:"type:float;not null"`

	// Cantidad de Habitaciones

	Rooms int `gorm:"type:int;not null"`

	// Amenities

	Parking bool `gorm:"type:bool;not null"`
	Pool    bool `gorm:"type:bool;not null"`
	Wifi    bool `gorm:"type:bool;not null"`
	Air     bool `gorm:"type:bool;not null"`
	Gym     bool `gorm:"type:bool;not null"`
	Spa     bool `gorm:"type:bool;not null"`

	Images Images `gorm:"foreignKey:HotelId"`
}

type Hotels []Hotel
