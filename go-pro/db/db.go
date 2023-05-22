package db

import (
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
	DBPass := ""
	//DBPass := os.Getenv("MVC_DB_PASS")
	DBHost := "localhost"
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	userClient.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.

	// AUTOMIGRATE recibe una estructura de GO como parametro
	// el GORM creo tablas dentro de la base de datos nuestra.

	// las estructuras de las tablas que crea estan en la carpeta "model"
	db.AutoMigrate(&model.User{}) // se basa en la estructura de model/user.go

	log.Info("Finishing Migration Database Tables")
}
