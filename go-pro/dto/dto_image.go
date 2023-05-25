package dto

type ImageDto struct {
	Id      int    `json:"id"`
	Url     string `json:"url"`
	HotelId int    `json:"hotel_id,omitempty"`
}

type ImagesDto []ImageDto
