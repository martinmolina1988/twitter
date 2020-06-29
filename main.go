package main

import (
	"github.com/martinmolina1988/twitter/handlers"
	"log"

	"github.com/martinmolina1988/twitter/bd"
)

func main() {
	if bd.ChequeoConexion() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}
