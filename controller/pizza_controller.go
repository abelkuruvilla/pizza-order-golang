package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	utils "pizza-delivery/apputil"
	"pizza-delivery/model"
	"pizza-delivery/store"
)

type Handler struct{}

func (h *Handler) BuyPizza(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	var req model.Request
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {

		return err, http.StatusBadRequest, fmt.Errorf("unable to decode JSON request body: %v", err)
	}
	utils.Debug("Incoming order", req)

	customer := req.GetCustomer()

	vegCount := req.CountVeg()
	meatCount := req.CountMeat()

	for i := 0; i < vegCount; i++ {
		store.StoreOrders.PlaceOrder(customer, &model.VegPizza{})
	}
	for i := 0; i < meatCount; i++ {
		store.StoreOrders.PlaceOrder(customer, &model.MeatPizza{})
	}
	return "Successfully placed order", http.StatusAccepted, nil
}

func (h *Handler) GetStatus(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	mobile := r.URL.Query().Get("mobile")
	ui, err := strconv.Atoi(mobile)
	if err != nil {
		return err, http.StatusBadRequest, err
	}
	data, err := store.StoreOrders.GetStatus(uint(ui))

	if err != nil {
		return err, http.StatusNotFound, err
	}
	utils.Info("Retrieved status ", data)
	return data, http.StatusFound, nil
}
