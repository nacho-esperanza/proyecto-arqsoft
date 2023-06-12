package clients

import (
	"go-pro/model"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetUserById(id int) model.User {
	var user model.User

	Db.Where("Id = ?", id).First(&user)
	log.Debug("User: ", user)

	return user
}

func GetUserByEmail(Email string) model.User {
	var user model.User

	Db.Where("email = ?", Email).First(&user)
	log.Debug("User: ", user)

	return user
}

func GetUsers() model.Users {
	var users model.Users

	// Recordar poner los preload para que traiga los datos de las tablas relacionadas

	Db.Find(&users)

	log.Debug("Users: ", users)

	return users
}

func InsertUser(user model.User) model.User {
	result := Db.Create(&user)

	if result.Error != nil {
		//TODO Manage Errors
		log.Error("")
	}
	log.Debug("User Created: ", user.Id)
	return user
}

// funcion que verifique si un usuario es admin si su email es admin@gmail.com

func IsAdmin(email string) bool {
	var user model.User
	Db.Where("email = ?", email).First(&user)

	if user.Email == "admin@gmail.com" {
		return true
	} else {
		return false
	}
}
