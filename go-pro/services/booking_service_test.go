package services

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-pro/dto"
	"go-pro/model"
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
	bookingService := bookingService{}

	bookingDto, err := bookingService.GetBookingById(1)

	// Verificar que el resultado sea el esperado
	assert.Equal(t, 1, bookingDto.Id)
	assert.Nil(t, err)
}
