package services

import (
	userCliente "go-pro/clients/user"

	"go-pro/dto"
	"go-pro/model"
	e "go-pro/utils/errors"
)

// El servicio procesa los datos y realiza operaciones de negocios utilizando los datos proporcionados por el model.

// interactua con el modelo

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDetailDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
	InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError)
}

var (
	UserService userServiceInterface
)

func init() {
	UserService = &userService{}
}

/*func (s *userService) GetUserById(id int) (dto.UserDetailDto, e.ApiError) {

	var user model.User = userCliente.GetUserById(id)
	var userDetailDto dto.UserDetailDto

	if user.Id == 0 {
		return userDetailDto, e.NewBadRequestApiError("user not found")
	}

	userDetailDto.Name = user.Name
	userDetailDto.Lastname = user.Lastname

	return userDetailDto, nil
}*/

func (s *userService) GetUserById(id int) (dto.UserDetailDto, e.ApiError) {
	user := userCliente.GetUserById(id)
	userDetailDto := dto.UserDetailDto{}

	if user.Id == 0 {
		return userDetailDto, e.NewBadRequestApiError("user not found")
	}

	userDetailDto.Name = user.Name
	userDetailDto.Lastname = user.LastName

	return userDetailDto, nil
}

/*func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {

	var users model.Users = userCliente.GetUsers()
	var usersDto dto.UsersDto

	for _, user := range users {
		var userDto dto.UserDto

		userDto.Id = user.Id
		userDto.Name = user.Name
		userDto.Lastname = user.Lastname
		userDto.Email = user.Email
		userDto.Password = user.Password

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}*/

func (s *userService) GetUsers() (dto.UsersDto, e.ApiError) {
	users := userCliente.GetUsers()
	usersDto := dto.UsersDto{}

	for _, user := range users {
		userDto := dto.UserDto{
			Id:       user.Id,
			Name:     user.Name,
			Lastname: user.LastName,
			Email:    user.Email,
			Password: user.Password,
		}

		usersDto = append(usersDto, userDto)
	}

	return usersDto, nil
}

func (s *userService) InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError) {

	var user model.User

	user.Name = userDto.Name
	user.LastName = userDto.Lastname
	user.Email = userDto.Email
	user.Password = userDto.Password

	user = userCliente.InsertUser(user)

	userDto.Id = user.Id

	return userDto, nil
}
