package db

import (
	addressClient "mvc-go/clients/address"
	telephoneClient "mvc-go/clients/telephone"
	userClient "mvc-go/clients/user"

	"mvc-go/model"

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
	DBName := "uccarqsoft"
	DBUser := "root"
	DBPass := "1234"
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
	addressClient.Db = db
	telephoneClient.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.

	// AUTOMIGRATE recibe una estructura de GO como parametro
	// el GORM creo tablas dentro de la base de datos nuestra.

	// las estructuras de las tablas que crea estan en la carpeta "model"
	db.AutoMigrate(&model.User{})      // se basa en la estructura de model/user.go
	db.AutoMigrate(&model.Address{})   // se basa en la estructura de model/adrdress.go
	db.AutoMigrate(&model.Telephone{}) // se basa en la estructura de model/telephone.go

	log.Info("Finishing Migration Database Tables")
}
