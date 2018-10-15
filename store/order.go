package store

import (
	"fmt"
	"os"
	"pizza-delivery/apputil"
	"pizza-delivery/bootstrapper"
	"pizza-delivery/model"
	"time"
)

//Store is used to define order customer with pizza
type Store struct {
	orders OrderMap
}

func (s *Store) PlaceOrder(customer *model.Customer, pizza PizzaInterface) {

	cs, err := s.FindCustomer(customer.Mobile)
	if err == nil {
		customer = cs
	}
	s.orders[customer] = append(s.orders[customer], pizza)

	//Pizza Ordered Time
	ot := time.Now().In(time.Local)
	go func(customer *model.Customer, pizza PizzaInterface, ot time.Time) {
		apputil.Info(fmt.Sprintf("Cooking Pizza for customer %s", customer.Name))
		c := make(chan model.PizzaStage)
		go pizza.Cook(c)
		for {
			select {
			case <-c:
				notifyCustomer(customer, pizza, ot)
			}
		}
	}(customer, pizza, ot)

}

func (s *Store) GetStatus(mobile uint) (interface{}, error) {

	customer, err := s.FindCustomer(mobile)

	if err != nil {
		return nil, err
	}
	var op []string
	for _, orders := range s.orders[customer] {
		op = append(op, orders.GetStage().String())
	}
	return op, nil
}
func (s *Store) FindCustomer(mobile uint) (*model.Customer, error) {
	for k, _ := range s.orders {
		if k.Mobile == mobile {
			return k, nil
		}
	}
	return &model.Customer{}, fmt.Errorf("No customer found")
}

var StoreOrders *Store

type orderList []PizzaInterface

type OrderMap map[*model.Customer]orderList

func init() {
	StoreOrders = &Store{
		orders: make(OrderMap),
	}
}

func notifyCustomer(customer *model.Customer, pizza PizzaInterface, ot time.Time) {

	apputil.Info(fmt.Sprintf("Notifying completed for customer %s", customer.Name))

	//Finished Time
	ft := time.Now().In(time.Local)
	pizType := pizza.GetType().String()

	fo, err := os.OpenFile(bootstrapper.AppConfig.NotifyFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		apputil.Error("Cannot open notify file", err)
		return
	}

	defer fo.Close()

	st := fmt.Sprintf("\nTime Started: %s Time Finished: %s -> Completed %s of customer Name: %s, Email: %s, Mobile: %v", ot.Format(time.UnixDate), ft.Format(time.UnixDate), pizType, customer.Name, customer.Email, customer.Mobile)
	fo.WriteString(st)
}
