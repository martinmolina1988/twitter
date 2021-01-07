package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/martinmolina1988/twitter/middlew"
	"github.com/martinmolina1988/twitter/routers"
	"github.com/rs/cors"
)

/*Manejadores seteo mi puerto, el Hnadler y pongo a escuchar al servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/verTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerTweet))).Methods("GET")
	router.HandleFunc("/modificarEmail", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarEmail))).Methods("PUT")
	router.HandleFunc("/modificarPassword", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPassword))).Methods("PUT")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/mensajeUsuario", middlew.ChequeoBD(middlew.ValidoJWT(routers.MensajeUsuario))).Methods("PUT")
	router.HandleFunc("/respuesta", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboRespuesta))).Methods("POST")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoRTs", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoRTs))).Methods("GET")
	router.HandleFunc("/leoCantidadDenuncias", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoCantidadDenuncias))).Methods("GET")
	router.HandleFunc("/leoCantidadRespuestas", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoCantidadRespuestas))).Methods("GET")
	router.HandleFunc("/leoRespuestas", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoRespuestas))).Methods("GET")
	router.HandleFunc("/leoNotificaciones", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoNotificaciones))).Methods("GET")
	router.HandleFunc("/leoTweets", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/leoUltimoMensaje", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoUltimoMensaje))).Methods("GET")
	router.HandleFunc("/leoMensajes", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoMensajes))).Methods("GET")
	router.HandleFunc("/leoUsuariosMensajes", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoUsuariosMensajes))).Methods("GET")
	router.HandleFunc("/leoTweetsTotal", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweetsTotal))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")
	router.HandleFunc("/eliminarNotif", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarNotif))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/envioMensaje", middlew.ChequeoBD(middlew.ValidoJWT(routers.EnvioMensaje))).Methods("POST")
	router.HandleFunc("/altaDenuncia", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaDenuncia))).Methods("POST")
	router.HandleFunc("/altaNotif", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaNotif))).Methods("POST")
	router.HandleFunc("/altaNotifLike", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaNotifLike))).Methods("POST")
	router.HandleFunc("/altaRTconMensaje", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaRTconMensaje))).Methods("POST")
	router.HandleFunc("/altaRT", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaRT))).Methods("POST")
	router.HandleFunc("/reseteoNotif", middlew.ChequeoBD(middlew.ValidoJWT(routers.ReseteoNotif))).Methods("GET")
	router.HandleFunc("/aumentoNotif", middlew.ChequeoBD(middlew.ValidoJWT(routers.AumentoNotif))).Methods("GET")
	router.HandleFunc("/altaLike", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaLike))).Methods("GET")
	router.HandleFunc("/bajaRT", middlew.ChequeoBD(middlew.ValidoJWT(routers.BajaRT))).Methods("GET")
	router.HandleFunc("/bajaLike", middlew.ChequeoBD(middlew.ValidoJWT(routers.BajaLike))).Methods("GET")
	router.HandleFunc("/altaBloqueo", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaBloqueo))).Methods("POST")
	router.HandleFunc("/altaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/bajaBloqueo", middlew.ChequeoBD(middlew.ValidoJWT(routers.BajaBloqueo))).Methods("DELETE")
	router.HandleFunc("/consultaBloqueo", middlew.ChequeoBD(middlew.ValidoJWT(routers.ConsultaBloqueo))).Methods("GET")
	router.HandleFunc("/consultaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")
	router.HandleFunc("/consultaLike", middlew.ChequeoBD(middlew.ValidoJWT(routers.ConsultaLike))).Methods("GET")
	router.HandleFunc("/consultaRT", middlew.ChequeoBD(middlew.ValidoJWT(routers.ConsultaRT))).Methods("GET")
	router.HandleFunc("/leoUsuariosConRTs", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoUsuariosConRTs))).Methods("GET")
	router.HandleFunc("/listaUsuarios", middlew.ChequeoBD(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/listaUsuariosTotal", middlew.ChequeoBD(middlew.ValidoJWT(routers.ListaUsuariosTotal))).Methods("GET")
	router.HandleFunc("/leoTweetsSeguidores", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweetsSeguidores))).Methods("GET")
	router.HandleFunc("/leoDenunciasPorCantidad", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoDenunciasPorCantidad))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
