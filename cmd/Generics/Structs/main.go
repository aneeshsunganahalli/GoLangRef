package main

import "fmt"

type gasEngine struct {
	mpg     float32
	gallons float32
}

type electricEngine struct {
	mpkwh float32
	kwh   float32
}

type car[T gasEngine | electricEngine] struct {
	carMake  string
	carModel string
	engine   T
}

func main() {
	var gasCar = car[gasEngine]{
		carMake:  "Honda",
		carModel: "Civic",
		engine: gasEngine{
			mpg:     10,
			gallons: 20,
		},
	}

	var electricCar = car[electricEngine]{
		carMake:  "Tesla",
		carModel: "S",
		engine: electricEngine{
			mpkwh: 102,
			kwh:   10,
		},
	}

	fmt.Print(gasCar, electricCar)
}