package jwt

import (
	"time"

	"log"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/martinmolina1988/twitter/models"

	"github.com/joho/godotenv"
)

func load() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")

	}

	var MiClave string = (os.Getenv("MiClave"))
	return MiClave
}

/*GeneroJWT genera el encriptado con JWT */
func GeneroJWT(t models.Usuario) (string, error) {

	miClave := []byte(load())

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
