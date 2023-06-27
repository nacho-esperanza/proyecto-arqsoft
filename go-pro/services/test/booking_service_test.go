package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-pro/dto"
	"go-pro/model"
	"go-pro/services"
	e "go-pro/utils/errors"
	"testing"
	"time"
)

type mockBookingClient struct {
	mock.Mock
}

func (m *mockBookingClient) GetBookingById(id int) (dto.BookingDto, e.ApiError) {
	args := m.Called(id)
	return args.Get(0).(dto.BookingDto), args.Error(1).(e.ApiError)
}

// Resto del código de las funciones mock...

func TestGetBookingById(t *testing.T) {
	// Crear instancia del cliente mock y configurarlo
	mockClient := new(mockBookingClient)
	mockClient.On("GetBookingById", 1).Return(model.Booking{
		Id:         1,
		HotelId:    1,
		UserId:     1,
		StartDate:  time.Now(),
		EndDate:    time.Now(),
		TotalPrice: 1000,
	}, nil)

	// Crear instancia del servicio con el cliente mock como parámetro
	bookingService := &services.bookingService{} // Utiliza la implementación concreta, no la interfaz

	bookingDto, err := bookingService.GetBookingById(1)

	// Verificar que el resultado sea el esperado
	assert.Equal(t, 1, bookingDto.Id)
	assert.Nil(t, err)
}

func (m *mockBookingClient) GetBookings() (dto.BookingsDto, e.ApiError) {
	args := m.Called()
	return args.Get(0).(dto.BookingsDto), args.Error(1).(e.ApiError)
}

func TestGetBookings(t *testing.T) {
	// Crear instancia del cliente mock y configurarlo
	mockClient := new(mockBookingClient)
	mockClient.On("GetBookings").Return(dto.BookingsDto{
		Bookings: []dto.BookingDto{
			{
				Id:         1,
				HotelId:    1,
				UserId:     1,
				StartDate:  time.Now(),
				EndDate:    time.Now(),
				TotalPrice: 1000,
			},
			{
				Id:         2,
				HotelId:    2,
				UserId:     2,
				StartDate:  time.Now(),
				EndDate:    time.Now(),
				TotalPrice: 2000,
			},
		},
	}, nil)

	// Crear instancia del servicio con el cliente mock como parámetro
	bookingService := bookingService{}

	bookingsDto, err := bookingService.GetBookings()

	// Verificar que el resultado sea el esperado
	assert.Equal(t, 2, len(bookingsDto.Bookings))
	assert.Nil(t, err)
}

func (m *mockBookingClient) CreateBooking(bookingDto dto.BookingDto) (dto.BookingDto, e.ApiError) {
	args := m.Called(bookingDto)
	return args.Get(0).(dto.BookingDto), args.Error(1).(e.ApiError)
}

func TestCreateBooking(t *testing.T) {
	// Crear instancia del cliente mock y configurarlo
	mockClient := new(mockBookingClient)
	mockClient.On("CreateBooking", mock.Anything).Return(dto.BookingDto{
		Id:         1,
		HotelId:    1,
		UserId:     1,
		StartDate:  time.Now(),
		EndDate:    time.Now(),
		TotalPrice: 1000,
	}, nil)

	// Crear instancia del servicio con el cliente mock como parámetro
	bookingService := &services.BookingService{Client: mockClient}

	bookingDto, err := bookingService.CreateBooking(dto.BookingDto{
		Id:         1,
		HotelId:    1,
		UserId:     1,
		StartDate:  time.Now(),
		EndDate:    time.Now(),
		TotalPrice: 1000,
	})

	// Verificar que el resultado sea el esperado
	assert.Equal(t, 1, bookingDto.Id)
	assert.Nil(t, err)
}
