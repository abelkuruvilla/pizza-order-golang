package router

import (
	"github.com/gorilla/mux"
)

// InitRoutes registers all routes for the application.
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	//Set routes for Pizza buy
	router = SetPizzaRoutes(router)
	return router
}
