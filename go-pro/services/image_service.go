package services

/*
import (
	hotelCliente "go-pro/clients/hotel"
	imageCliente "go-pro/clients/image"
	"go-pro/dto"
	"go-pro/model"
	e "go-pro/utils/errors"
)

// El servicio procesa los datos y realiza operaciones de negocios utilizando los datos proporcionados por el model.

// interactua con el modelo

type imageService struct{}

type imageServiceInterface interface {
	AddImage(imageDto dto.ImageDto) (dto.ImageDto, e.ApiError)
}

var (
	ImageService imageServiceInterface
)

func init() {
	ImageService = &imageService{}
}

func (s *imageService) AddImage(imageDto dto.ImageDto) (dto.ImageDto, e.ApiError) {

	var image model.Image

	// Verificar si el hotel existe
	hotel := hotelCliente.GetHotelById(imageDto.HotelId)
	if hotel.Id == 0 {
		// El hotel no existe, devolver un error
		return dto.ImageDto{}, e.NewBadRequestApiError("Hotel not found")
	}

	image.Url = imageDto.Url
	image.HotelId = imageDto.HotelId

	image = imageCliente.AddImage(image)

	imageDto.Id = image.Id

	return imageDto, nil
}
*/
