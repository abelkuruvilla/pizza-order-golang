package model

import (
	"time"
)

const (
	//VEGGIE type denotes vegetable pizza
	VEGGIE PizzaType = iota //0
	//MEAT type denotes Non-veg pizza
	MEAT
)

const (
	//PREP denotes first stage Dough Prepare
	PREP PizzaStage = iota //0
	//BAKE denotes second stage Oven Bake
	BAKE
	//ART denotes third stage Topping Art
	ART
	//FINISH denotes finished Pizza Stage
	FINISH
)

type (
	//PizzaType is used to define types of Pizzas
	PizzaType int
	//PizzaStage defines different stages in pizza cooking
	PizzaStage int
	pizzaBase  struct {
		Stage PizzaStage
	}
	//VegPizza type is used to define Veg Pizza
	VegPizza struct {
		pizzaBase
	}
	//MeatPizza type is used to define Meat Pizza
	MeatPizza struct {
		pizzaBase
	}
)

//GetStage is used to return the current stage of pizza order
func (p *pizzaBase) GetStage() PizzaStage {
	return p.Stage
}

//Cook is used to emulate different stages of cooking
func (p *VegPizza) Cook(stage chan PizzaStage) {
	p.Stage = PREP
	time.Sleep(5 * time.Second)
	p.Stage = BAKE
	time.Sleep(3 * time.Second)
	p.Stage = ART
	time.Sleep(4 * time.Second)
	p.Stage = FINISH
	stage <- p.Stage
}

func (p *VegPizza) GetType() PizzaType {
	return VEGGIE
}

func (p *MeatPizza) GetType() PizzaType {
	return MEAT
}

//Cook is used to emulate different stages of cooking
func (p *MeatPizza) Cook(stage chan PizzaStage) {
	p.Stage = PREP
	time.Sleep(3 * time.Second)
	p.Stage = BAKE
	time.Sleep(7 * time.Second)
	p.Stage = ART
	time.Sleep(5 * time.Second)
	p.Stage = FINISH

	stage <- p.Stage
}

func (t PizzaType) String() string {
	var s string
	switch t {
	case MEAT:
		s = "Meat Pizza"
	case VEGGIE:
		s = "Vegitable Pizza"
	}
	return s
}

func (ps PizzaStage) String() string {
	var s string
	switch ps {
	case PREP:
		s = "Prepare Dough"
	case BAKE:
		s = "Oven Bake"
	case ART:
		s = "Topping Art"
	case FINISH:
		s = "Finished Making"
	}
	return s
}
