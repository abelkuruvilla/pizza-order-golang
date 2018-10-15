package router

import (
	"pizza-delivery/controller"

	"github.com/gorilla/mux"
)

func SetPizzaRoutes(router *mux.Router) *mux.Router {

	h := controller.Handler{}
	router.Handle("/buy_pizza", responseHandler(h.BuyPizza)).Methods("POST")
	router.Handle("/status", responseHandler(h.GetStatus)).Methods("GET")
	return router
}
