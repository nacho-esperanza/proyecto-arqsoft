package services

import (
	hotelCliente "go-pro/clients/hotel"

	"go-pro/dto"
	"go-pro/model"
	e "go-pro/utils/errors"
)

// El servicio procesa los datos y realiza operaciones de negocios utilizando los datos proporcionados por el model.

// interactua con el modelo

type hotelService struct{}

type hotelServiceInterface interface {
	GetHotelById(id int) (dto.HotelDto, e.ApiError)
	GetHotels() (dto.HotelsDto, e.ApiError)
	InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError)
}

var (
	HotelService hotelServiceInterface
)

func init() {
	HotelService = &hotelService{}
}

func (s *hotelService) GetHotelById(id int) (dto.HotelDto, e.ApiError) {
	hotel := hotelCliente.GetHotelById(id)
	hotelDto := dto.HotelDto{}

	if hotel.Id == 0 {
		return hotelDto, e.NewBadRequestApiError("hotel not found")
	}

	hotelDto.Name = hotel.Name
	hotelDto.Adress = hotel.Adress
	hotelDto.City = hotel.City
	hotelDto.Stars = hotel.Stars
	hotelDto.Description = hotel.Description
	hotelDto.Price = hotel.Price

	// Amenities
	hotelDto.Parking = hotel.Parking
	hotelDto.Pool = hotel.Pool
	hotelDto.Wifi = hotel.Wifi
	hotelDto.Air = hotel.Air
	hotelDto.Gym = hotel.Gym
	hotelDto.Spa = hotel.Spa

	return hotelDto, nil

}

// Funcion que muestra TODOS los hoteles
func (s *hotelService) GetHotels() (dto.HotelsDto, e.ApiError) {
	hotels := hotelCliente.GetHotels()
	hotelsDto := dto.HotelsDto{}

	for _, hotel := range hotels {
		hotelDto := dto.HotelDto{}

		hotelDto.Id = hotel.Id
		hotelDto.Name = hotel.Name
		hotelDto.Adress = hotel.Adress
		hotelDto.City = hotel.City
		hotelDto.Stars = hotel.Stars
		hotelDto.Description = hotel.Description
		hotelDto.Price = hotel.Price

		// Amenities
		hotelDto.Parking = hotel.Parking
		hotelDto.Pool = hotel.Pool
		hotelDto.Wifi = hotel.Wifi
		hotelDto.Air = hotel.Air
		hotelDto.Gym = hotel.Gym
		hotelDto.Spa = hotel.Spa

		hotelsDto = append(hotelsDto, hotelDto)
	}

	return hotelsDto, nil
}

// Funcion que inserta un hotel

func (s *hotelService) InsertHotel(hotelDto dto.HotelDto) (dto.HotelDto, e.ApiError) {

	var hotel model.Hotel

	hotel.Name = hotelDto.Name
	hotel.Adress = hotelDto.Adress
	hotel.City = hotelDto.City
	hotel.Stars = hotelDto.Stars
	hotel.Description = hotelDto.Description
	hotel.Price = hotelDto.Price

	// Amenities
	hotel.Parking = hotelDto.Parking
	hotel.Pool = hotelDto.Pool
	hotel.Wifi = hotelDto.Wifi
	hotel.Air = hotelDto.Air
	hotel.Gym = hotelDto.Gym
	hotel.Spa = hotelDto.Spa

	hotel = hotelCliente.InsertHotel(hotel)

	hotelDto.Id = hotel.Id

	return hotelDto, nil
}
