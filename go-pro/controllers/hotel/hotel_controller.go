package hotelController

import (
	"go-pro/dto"
	service "go-pro/services"
	"net/http"
	"strconv"

	e "go-pro/utils/errors"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetHotelById(c *gin.Context) {
	log.Debug("Hotel id to load: " + c.Param("id"))

	id, _ := strconv.Atoi(c.Param("id"))
	var hotelDto dto.HotelDto

	hotelDto, err := service.HotelService.GetHotelById(id) // se utiliza el userDto referenciandose del id

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, hotelDto)
}

func GetHotels(c *gin.Context) {
	var hotelsDto dto.HotelsDto
	hotelsDto, err := service.HotelService.GetHotels()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, hotelsDto)
}

func HotelInsert(c *gin.Context) {
	var hotelDto dto.HotelDto
	err := c.BindJSON(&hotelDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	hotelDto, er := service.HotelService.InsertHotel(hotelDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, hotelDto)
}

func AddHotelImage(c *gin.Context) {
	// Obtener el ID del hotel de los parámetros de la URL
	hotelId, err := strconv.Atoi(c.Param("hotelId"))
	if err != nil {
		apiError := e.NewBadRequestApiError("Invalid hotel Id")
		c.JSON(apiError.Status(), apiError)
		return
	}

	// Obtener los datos de la imagen del cuerpo de la solicitud
	var imageDto dto.ImageDto
	if err := c.ShouldBindJSON(&imageDto); err != nil {
		apiError := e.NewBadRequestApiError("Invalid image data")
		c.JSON(apiError.Status(), apiError)
		return
	}

	// Agregar la imagen al hotel utilizando el servicio correspondiente
	image, apiError := service.HotelService.AddHotelImage(hotelId, imageDto)
	if apiError != nil {
		c.JSON(apiError.Status(), apiError)
		return
	}

	c.JSON(http.StatusOK, image)

	/*var imageDto dto.ImageDto
	err := c.BindJSON(&imageDto)

	// Error Parsing json param
	if err != nil {
		log.Error(err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	imageDto, er := service.HotelService.AddHotelImage(id, imageDto)
	// Error del Insert
	if er != nil {
		c.JSON(er.Status(), er)
		return
	}

	c.JSON(http.StatusCreated, imageDto)
	*/
}

func GetHotelImagesById(c *gin.Context) {

	log.Debug("Hotel id to load: " + c.Param("hotelId"))

	hotelId := c.Param("hotelId")
	/*
		hotelIdStr := c.Param("hotelId")
		hotelId, _ := strconv.Atoi(hotelIdStr)
		var imagesDto dto.ImagesDto*/

	imagesDto, err := service.HotelService.GetHotelImagesById(hotelId) // se utiliza el userDto referenciandose del id

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, imagesDto)

	/*
		// Obtener el ID del hotel de los parámetros de la URL
		hotelId, err := strconv.Atoi(c.Param("hotelId"))
		if err != nil {
			apiError := e.NewBadRequestApiError("Invalid hotel Id")
			c.JSON(apiError.Status(), apiError)
			return
		}

		// Llamar al servicio para obtener las imágenes del hotel
		imagesDto, apiError := service.HotelService.GetHotelImagesById(hotelId)
		if apiError != nil {
			c.JSON(apiError.Status(), apiError)
			return
		}

		// Devolver la respuesta con los DTO de las imágenes
		c.JSON(http.StatusOK, imagesDto)
	*/
}

func GetHotelImages(c *gin.Context) {
	var imagesDto dto.ImagesDto
	imagesDto, err := service.HotelService.GetHotelImages()

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusOK, imagesDto)
}
