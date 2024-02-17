package db

import (
	bookingClient "go-pro/clients/booking"
	hotelClient "go-pro/clients/hotel"
	imageClient "go-pro/clients/image"
	userClient "go-pro/clients/user"

	"go-pro/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Paramters
	DBName := "go_booking_db"
	DBUser := "root"
	DBPass := "root"

	//DBPass := os.Getenv("MVC_DB_PASS")
	DBHost := "localhost"

	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3307)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	userClient.Db = db
	hotelClient.Db = db
	imageClient.Db = db
	bookingClient.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.

	// AUTOMIGRATE recibe una estructura de GO como parametro
	// el GORM creo tablas dentro de la base de datos nuestra.

	// las estructuras de las tablas que crea estan en la carpeta "model"
	db.AutoMigrate(&model.User{})    // se basa en la estructura de model/user.go
	db.AutoMigrate(&model.Hotel{})   // se basa en la estructura de model/hotel.go
	db.AutoMigrate(&model.Image{})   // se basa en la estructura de model/image.go
	db.AutoMigrate(&model.Booking{}) // se basa en la estructura de model/booking.go

	log.Info("Finishing Migration Database Tables")
}
