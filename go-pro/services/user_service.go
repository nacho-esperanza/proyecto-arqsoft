package services

import (
	userCliente "go-pro/clients/user"

	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"

	"go-pro/dto"
	"go-pro/model"
	e "go-pro/utils/errors"

	"crypto/md5"
	"encoding/hex"
)

// El servicio procesa los datos y realiza operaciones de negocios utilizando los datos proporcionados por el model.

// interactua con el modelo

type userService struct{}

type userServiceInterface interface {
	GetUserById(id int) (dto.UserDetailDto, e.ApiError)
	GetUsers() (dto.UsersDto, e.ApiError)
	InsertUser(userDto dto.UserDto) (dto.UserDto, e.ApiError)
	LoginUser(loginDto dto.LoginDto) (dto.TokenDto, e.ApiError)
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
	userDetailDto.Email = user.Email

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

	log.Debug(userDto) // Imprimir en consola el userDto ingresado
	var user model.User = userCliente.GetUserByEmail(userDto.Email)

	// PREGUNTAR A EMI SI HACE FALTA EL TOKEN O NO EN ESTA FUNCION

	if user.Id == 0 { // Si el usuario no existe, crearlo

		user.Name = userDto.Name
		user.LastName = userDto.Lastname
		user.Email = userDto.Email
		user.Password = hashMD5(userDto.Password) // Hashear la contraseña usando la función hashMD5

		user = userCliente.InsertUser(user)

		userDto.Id = user.Id

		return userDto, nil
	} else { // Si el usuario ya existe, devolver error
		return userDto, e.NewBadRequestApiError("email already registered")
	}
}

// FUNCION HASHMD5 PARA HASHEAR CONTRASEÑA Y QUE NO SE MUESTRE TAL CUAL EN LA BASE DE DATOS

func hashMD5(password string) string {
	hasher := md5.New()
	hasher.Write([]byte(password))
	hashedBytes := hasher.Sum(nil)
	hashedPassword := hex.EncodeToString(hashedBytes)
	return hashedPassword
}

// FUNCION PARA LOGUEAR USUARIO

var jwtKey = []byte("secret_key")

func (s *userService) LoginUser(loginDto dto.LoginDto) (dto.TokenDto, e.ApiError) {

	log.Debug(loginDto) // Imprimir en consola el loginDto
	var user model.User = userCliente.GetUserByEmail(loginDto.Email)

	var tokenDto dto.TokenDto

	if user.Id == 0 {
		return tokenDto, e.NewBadRequestApiError("user not found")
	}

	/*
		//pasamos password como slice de bytes
		//hashea con md5.sum
		var pswMd5 = md5.Sum([]byte(loginDto.Password))

		//convertir a cadena hexadecimal
		pswMd5String := hex.EncodeToString(pswMd5[:])
	*/

	//hashear la contraseña
	var pswMd5 = hashMD5(loginDto.Password)

	log.Debug(pswMd5)        // Imprimir en consola el pswMd5
	log.Debug(user.Password) // Imprimir en consola la password del usuario

	//comparar contraseñas

	// hacer parse token
	if pswMd5 == user.Password {
		//se firma el token para verificar autenticidad
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id_user": user.Id,
		})
		tokenString, _ := token.SignedString(jwtKey)
		tokenDto.Token = tokenString
		tokenDto.IdUser = user.Id

		return tokenDto, nil
	} else {
		return tokenDto, e.NewBadRequestApiError("contraseña incorrecta")
	}
	/*
		if pswMd5String == user.Password {
			//se firma el token para verificar autenticidad
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"id_user": user.Id,
			})
			tokenString, _ := token.SignedString(jwtKey)
			tokenDto.Token = tokenString
			tokenDto.IdUser = user.Id

			return tokenDto, nil
		} else {
			return tokenDto, e.NewBadRequestApiError("contraseña incorrecta")
		}
	*/

}
