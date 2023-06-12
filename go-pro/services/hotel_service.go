package services

import (
	hotelCliente "go-pro/clients/hotel"
	imageCliente "go-pro/clients/image"

	"strconv"

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
	AddHotelImage(hotelId int, imageDto dto.ImageDto) (dto.ImageDto, e.ApiError)
	GetHotelImagesById(hotelId string) (dto.ImagesDto, e.ApiError)
	GetHotelImages() (dto.ImagesDto, e.ApiError)
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

	hotelDto.Id = hotel.Id
	hotelDto.Name = hotel.Name
	hotelDto.Adress = hotel.Adress
	hotelDto.City = hotel.City
	hotelDto.Stars = hotel.Stars
	hotelDto.Description = hotel.Description
	hotelDto.Price = hotel.Price

	// Habitaciones disponibles
	hotelDto.Rooms = hotel.Rooms

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

		// Habitaciones disponibles
		hotelDto.Rooms = hotel.Rooms

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

	// Habitaciones disponibles
	hotel.Rooms = hotelDto.Rooms

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

// Funcion que inserta una imagen

func (s *hotelService) AddHotelImage(hotelId int, imageDto dto.ImageDto) (dto.ImageDto, e.ApiError) {

	hotel := hotelCliente.GetHotelById(hotelId)
	if hotel.Id == 0 {
		return dto.ImageDto{}, e.NewBadRequestApiError("hotel not found")
	}

	image := model.Image{
		Url:     imageDto.Url,
		HotelId: hotelId,
	}

	image = hotelCliente.AddHotelImage(image)

	imageDto.Id = image.Id

	return imageDto, nil
	/*
		var image model.Image

		image.HotelId = hotelId
		image.Url = imageDto.Url

		image = hotelCliente.AddHotelImage(image)

		imageDto.Id = image.Id

		return imageDto, nil
	*/
}

// Funcion que muestra las imagenes de un hotel

func (s *hotelService) GetHotelImagesById(hotelId string) (dto.ImagesDto, e.ApiError) {

	id, err := strconv.Atoi(hotelId)
	if err != nil {
		// Manejar el error de conversión si hotelId no es un número válido
		return dto.ImagesDto{}, e.NewBadRequestApiError("Invalid hotelId")
	}

	// Verificar si el hotel existe
	hotel := hotelCliente.GetHotelById(id)
	if hotel.Id == 0 {
		return dto.ImagesDto{}, e.NewBadRequestApiError("hotel not found")
	}

	// Obtener las imágenes del hotel
	images := imageCliente.GetHotelImagesById(id)

	// Convertir las imágenes al DTO correspondiente
	imagesDto := dto.ImagesDto{}
	for _, image := range images {
		imageDto := dto.ImageDto{
			Id:      image.Id,
			Url:     image.Url,
			HotelId: image.HotelId,
		}
		imagesDto = append(imagesDto, imageDto)
	}

	return imagesDto, nil
}

// Funcion que muestra TODAS las imagenes de TODOS los hoteles
func (s *hotelService) GetHotelImages() (dto.ImagesDto, e.ApiError) {

	images := imageCliente.GetHotelImages()

	imagesDto := dto.ImagesDto{}

	for _, image := range images {
		imageDto := dto.ImageDto{}

		imageDto.Id = image.Id
		imageDto.Url = image.Url
		imageDto.HotelId = image.HotelId

		imagesDto = append(imagesDto, imageDto)
	}

	return imagesDto, nil
}
