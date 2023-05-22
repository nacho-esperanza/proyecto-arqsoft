package model

type User struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(50);not null"`
	LastName string `gorm:"type:varchar(50);not null"`
	Email    string `gorm:"type:varchar(100);not null,unique"`
	Password string `gorm:"type:varchar(255);not null"`
}

type Users []User
