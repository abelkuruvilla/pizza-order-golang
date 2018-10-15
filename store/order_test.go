package store

import (
	"pizza-delivery/model"
	"testing"
)

//Cannot run this successfully since the notify customer requires bootstrapper package
func TestPlaceOrder(t *testing.T) {

	testStore := Store{
		orders: make(OrderMap),
	}
	testCust := &model.Customer{Name: "Abel Kuruvilla", Email: "abel@qwe.com", Mobile: 8089352216}
	testVegPizza := &model.VegPizza{}
	testStore.PlaceOrder(testCust, testVegPizza)
	if !doesArrayContainElem(testStore.orders[testCust], testVegPizza) {
		t.Errorf("Failed in placing Order. Order not in array")
	}
}

func doesArrayContainElem(l orderList, p PizzaInterface) bool {
	for _, elem := range l {
		if elem == p {
			return true
		}
	}
	return false
}
