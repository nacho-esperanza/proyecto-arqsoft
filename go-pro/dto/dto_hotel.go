package dto

type HotelDto struct {
	Id          int     `json:"id"`
	Name        string  `json:"name"`
	Adress      string  `json:"adress"`
	City        string  `json:"city"`
	Stars       int     `json:"stars"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`

	// Cantidad de Habitaciones

	Rooms int `json:"rooms"`

	// Amenities

	Parking bool `json:"parking"`
	Pool    bool `json:"pool"`
	Wifi    bool `json:"wifi"`
	Air     bool `json:"air"`
	Gym     bool `json:"gym"`
	Spa     bool `json:"spa"`

	Images []ImageDto `json:"images"`
}

type HotelsDto []HotelDto
