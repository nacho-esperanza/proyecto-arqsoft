package model

// many2many: para poner MUCHOS A MUCHOS.
// despues del gorm:""

// no hace falta escribir la RELACION DOBLE para crear la base de datos.
// pero si se necesita al menos 1 de las 2 relaciones cuando hay MUCHOS A MUCHOS.

type User struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(350);not null"`
	LastName string `gorm:"type:varchar(250);not null"`
	UserName string `gorm:"type:varchar(150);not null;unique"`
	Password string `gorm:"type:varchar(150);not null"`

	Address   Address `gorm:"foreignkey:AddressId"`
	AddressId int

	Telephones Telephones `gorm:"foreignKey:UserId"`
}

type Users []User
