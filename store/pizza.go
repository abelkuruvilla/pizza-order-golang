package store

import "pizza-delivery/model"

//PizzaInterface - An interface to recieve pizza
type PizzaInterface interface {
	Cook(chan model.PizzaStage)
	GetStage() model.PizzaStage
	GetType() model.PizzaType
}
