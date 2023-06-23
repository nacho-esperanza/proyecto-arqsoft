package services

import (
	bookingCliente "go-pro/clients/booking"
	hotelCliente "go-pro/clients/hotel"
	userCliente "go-pro/clients/user"

	"fmt"
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
	GetBookingsByUserId(userId int) (dto.BookingsDto, e.ApiError)
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

	// Verificar si el hotel tiene habitaciones disponibles
	bookingsCount := bookingCliente.GetBookingsByDateRange(hotel.Id, startDate, endDate)
	availableRooms := hotel.Rooms - bookingsCount

	fmt.Println(fmt.Sprintf("Total rooms: %d vs Available: %d vs Bookings: %d", hotel.Rooms, availableRooms, bookingsCount))

	// Verificar si hay habitaciones disponibles para cada día en el rango de fechas
	currentDate := startDate
	for currentDate.Before(endDate) {
		bookingsCount := bookingCliente.GetBookingsByDateRange(hotel.Id, currentDate, currentDate)
		availableRooms = hotel.Rooms - bookingsCount

		if availableRooms <= 0 {
			return dto.BookingDto{}, e.NewBadRequestApiError("Hotel has no available rooms for the specified dates")
		}

		currentDate = currentDate.AddDate(0, 0, 1) // Avanzar al siguiente día
	}

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

func (s *bookingService) GetBookingsByUserId(userId int) (dto.BookingsDto, e.ApiError) {
	bookings := bookingCliente.GetBookingsByUserId(userId)

	user := userCliente.GetUserById(userId)
	if user.Id == 0 {
		return dto.BookingsDto{}, e.NewBadRequestApiError("User not found")
	}

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
