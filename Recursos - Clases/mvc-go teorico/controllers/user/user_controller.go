package userController

import (
	"mvc-go/dto"
	service "mvc-go/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// El controller recibe un Get de un User ID en donde verifica que ex
// Se utiliza el objeto Dto.User para transferir los datos ingresados por el usuario en view al service,
// consulta en el model y si los encuentra devuelve un 200 OK con el resultado.

// DTO TRANSPORTA DATOS SOLAMENTE

func GetUserById(c *gin.Context) {
	log.Debug("User id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var userDto dto.UserDetailDto

	userDto, err := service.UserService.GetUserById(id) // se utiliza el userDto referenciandose del id

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, userDto)
}

func GetUsers(c *gin.Context) {
	var usersDto dto.UsersDto
	usersDto, err := service.UserService.GetUsers()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, usersDto)
}

func UserInsert(c *gin.Context) {
	var userDto dto.UserDto
	err := c.BindJSON(&userDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	userDto, er := service.UserService.InsertUser(userDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, userDto)
}

func AddUserTelephone(c *gin.Context) {

	log.Debug("Adding Telephone to user: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id"))

	var telephoneDto dto.TelephoneDto
	err := c.BindJSON(&telephoneDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	telephoneDto.UserId = id

	var userDto dto.UserDetailDto

	userDto, er := service.UserService.AddUserTelephone(telephoneDto)
	//Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, userDto)
}
