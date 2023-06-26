package imageController

/*
import (
	"go-pro/dto"
	service "go-pro/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)
*/
// El controller recibe un Get de un Image ID en donde verifica que ex
// Se utiliza el objeto Dto.Image para transferir los datos ingresados por el usuario en view al service,
// consulta en el model y si los encuentra devuelve un 200 OK con el resultado.

// DTO TRANSPORTA DATOS SOLAMENTE

/*

func AddImage(c *gin.Context) {
	log.Debug("Adding Image to hotel: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))

	var imageDto dto.ImageDto
	err := c.BindJSON(&imageDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	imageDto.HotelId = id

	var hotelDto dto.HotelDto

	hotelDto, err = service.HotelService.AddImage(imageDto)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusCreated, hotelDto)

}

*/
