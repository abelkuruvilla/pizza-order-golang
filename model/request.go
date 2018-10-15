package model

//Request structure is used to define incoming POST request
type Request struct {
	Veg  int `json:"veg"`
	Meat int `json:"meat"`
	Customer
}

func (o Request) GetCustomer() *Customer {
	var c Customer
	c.Email = o.Email
	c.Mobile = o.Mobile
	c.Name = o.Name

	return &c
}

func (o Request) CountVeg() int {
	return o.Veg
}

func (o Request) CountMeat() int {
	return o.Meat
}
