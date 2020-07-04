package main

import (
	"log"

	"github.com/martinmolina1988/twitter/bd"
	"github.com/martinmolina1988/twitter/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}
