package main

import (
	"go-pro/app"
	"go-pro/db"
)

func main() {
	db.StartDbEngine()
	app.StartRoute()
}

/*Segun el profe:
1. Hacer el Client
2. Hacer el Servicio
3. Hacer el Controlador
*/

/*
BUENA PRACTICA:
	- Empezar haciendo el diagrama de clases.
*/
