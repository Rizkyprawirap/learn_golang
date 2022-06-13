package app

import (
	"github.com/rizkyprawirap/Toko/app/controller"
	"github.com/gorilla/mux"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controller.Home).Methods("GET")
}
