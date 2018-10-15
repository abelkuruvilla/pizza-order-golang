package model

import "testing"

func TestVegGetType(t *testing.T) {
	cases := []VegPizza{
		VegPizza{},
		VegPizza{},
		VegPizza{},
		VegPizza{},
	}
	for i := 0; i < 4; i++ {
		cases[i].Stage = PizzaStage(i)
	}

	for _, test := range cases {
		if ty := test.GetType().String(); ty != "Vegitable Pizza" {
			t.Errorf("Didnot get Expected Type. Got: %s . Want: %s", ty, "Vegitable Pizza")
		}
	}
}

func TestMeatGetType(t *testing.T) {
	cases := []MeatPizza{
		MeatPizza{},
		MeatPizza{},
		MeatPizza{},
		MeatPizza{},
	}
	for i := 0; i < 4; i++ {
		cases[i].Stage = PizzaStage(i)
	}

	for _, test := range cases {
		if ty := test.GetType().String(); ty != "Meat Pizza" {
			t.Errorf("Didnot get Expected Type. Got: %s . Want: %s", ty, "Meat Pizza")
		}
	}
}
