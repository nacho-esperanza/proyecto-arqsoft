package services

import (
	bookingCliente "go-pro/clients/booking"
	hotelCliente "go-pro/clients/hotel"
	userCliente "go-pro/clients/user"

	"go-pro/dto"
	"go-pro/model"
	e "go-pro/utils/errors"
	"time"
)

// El servicio procesa los datos y realiza operaciones de negocios utilizando los datos proporcionados por el model.

// interactua con el modelo

type bookingService struct{}

type bookingServiceInterface interface {
	GetBookingById(id int) (dto.BookingDto, e.ApiError)
	GetBookings() (dto.BookingsDto, e.ApiError)
	CreateBooking(bookingDto dto.BookingDto) (dto.BookingDto, e.ApiError)
}

var (
	BookingService bookingServiceInterface
)

func init() {
	BookingService = &bookingService{}
}

func (s *bookingService) GetBookingById(id int) (dto.BookingDto, e.ApiError) {
	booking := bookingCliente.GetBookingById(id)

	bookingDto := dto.BookingDto{}

	if booking.Id == 0 {
		return bookingDto, e.NewBadRequestApiError("booking not found")
	}

	bookingDto.Id = booking.Id
	bookingDto.HotelId = booking.HotelId
	bookingDto.UserId = booking.UserId
	bookingDto.StartDate = booking.StartDate
	bookingDto.EndDate = booking.EndDate
	bookingDto.TotalPrice = booking.TotalPrice

	return bookingDto, nil
}

func (s *bookingService) CreateBooking(bookingDto dto.BookingDto) (dto.BookingDto, e.ApiError) {
	// Verificar si el hotel existe
	hotel := hotelCliente.GetHotelById(bookingDto.HotelId)
	if hotel.Id == 0 {
		return dto.BookingDto{}, e.NewBadRequestApiError("Hotel not found")
	}

	// Verificar si el usuario existe
	user := userCliente.GetUserById(bookingDto.UserId)
	if user.Id == 0 {
		return dto.BookingDto{}, e.NewBadRequestApiError("User not found")
	}

	// Calcular el TotalPrice a partir del precio del hotel y las fechas
	startDate, err := time.Parse("2006-01-02", bookingDto.StartDate.Format("2006-01-02"))
	if err != nil {
		return dto.BookingDto{}, e.NewBadRequestApiError("Invalid start date format")
	}

	endDate, err := time.Parse("2006-01-02", bookingDto.EndDate.Format("2006-01-02"))
	if err != nil {
		return dto.BookingDto{}, e.NewBadRequestApiError("Invalid end date format")
	}

	// totalPrice := int(hotel.Price * float64(endDate.Sub(startDate).Hours()/24))
	totalPrice := float64(hotel.Price) * endDate.Sub(startDate).Hours() / 24

	// Crear la entidad Booking
	booking := model.Booking{
		HotelId:    bookingDto.HotelId,
		UserId:     bookingDto.UserId,
		StartDate:  startDate,
		EndDate:    endDate,
		TotalPrice: totalPrice,
	}

	// Guardar el booking en la base de datos
	createdBooking := bookingCliente.CreateBooking(booking)

	/*

		// PREGUNTAR ESTO

	*/
	/*
		// Obtener la fecha actual
		now := time.Now()

		// Verificar si la fecha actual está dentro del rango de fechas de la reserva
		IsActive := now.After(createdBooking.StartDate) && now.Before(createdBooking.EndDate)

		// Actualizar el número de habitaciones disponibles en el hotel si la reserva está activa
		if createdBooking.Id != 0 && IsActive {
			hotel.Rooms -= 1
			hotelCliente.UpdateHotel(hotel)
		}
	*/

	/*

		// Obtener la fecha actual
		now := time.Now()

		// Creo booleano IsActive

		// Inicializar IsActive en false por defecto
		IsActive := false

		// Verificar si la fecha actual está dentro del rango de fechas de la reserva
		if now.After(createdBooking.StartDate) && now.Before(createdBooking.EndDate) {
			// La reserva está activa si la fecha actual está después de la fecha de inicio
			// y antes de la fecha de fin
			IsActive = true
		} else {
			// La reserva no está activa si la fecha actual está fuera del rango de fechas
			IsActive = false
		}

		// Actualizar el número de habitaciones disponibles en el hotel
		if createdBooking.Id != 0 && IsActive {

			// Restar 1 a las habitaciones disponibles si la reserva está activa
			hotel.Rooms -= 1

			// Actualizar el hotel en la base de datos
			hotelCliente.UpdateHotel(hotel)
		} else {
			// Sumar 1 a las habitaciones disponibles si la reserva deja de estar activa
			hotel.Rooms += 1

			// Actualizar el hotel en la base de datos
			hotelCliente.UpdateHotel(hotel)
		}

	*/

	// Crear el DTO de respuesta
	responseDto := dto.BookingDto{
		Id:         createdBooking.Id,
		HotelId:    createdBooking.HotelId,
		UserId:     createdBooking.UserId,
		StartDate:  createdBooking.StartDate,
		EndDate:    createdBooking.EndDate,
		TotalPrice: createdBooking.TotalPrice,
	}

	return responseDto, nil
}

func (s *bookingService) GetBookings() (dto.BookingsDto, e.ApiError) {
	bookings := bookingCliente.GetBookings()

	bookingsDto := dto.BookingsDto{}

	for _, booking := range bookings {
		bookingDto := dto.BookingDto{
			Id:         booking.Id,
			HotelId:    booking.HotelId,
			UserId:     booking.UserId,
			StartDate:  booking.StartDate,
			EndDate:    booking.EndDate,
			TotalPrice: booking.TotalPrice,
		}
		bookingsDto.Bookings = append(bookingsDto.Bookings, bookingDto)
	}

	return bookingsDto, nil
}
