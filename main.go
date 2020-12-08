package main

import (
	"fmt"
	"log"

	socketio "github.com/googollee/go-socket.io"
	"github.com/martinmolina1988/twitter/bd"
	"github.com/martinmolina1988/twitter/handlers"
)

func main() {

	server, err := socketio.NewServer(nil)

	if err != nil {
		log.Fatal(err)
	}

	//sockets
	server.OnConnect("/", func(so socketio.Conn) error {
		so.SetContext("")
		fmt.Println("nuevo usuario conectado")
		return nil
	})

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}
