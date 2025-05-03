package main

import "fmt"

type gasEngine struct {		// Reusable
	mpg     uint8
	gallons uint8
}

type electricEngine struct {
	mpkwh uint8
	kwh uint8
}
func (e electricEngine) milesLeft() uint8{
	return e.kwh * e.mpkwh
}

// Methods correlated to structs
func (e gasEngine) milesLeft() uint8 { 		// (e gasEgine) is how you assign the function to the gasEngine struct
	return e.gallons * e.mpg
}

type engine interface{
	milesLeft() uint8
}

func canMakeIt(e gasEngine, miles uint8) {		// Can only take gas engine param, but if I have other similar engines, need to make interface
	if miles <= e.milesLeft(){
		fmt.Printf("Can make it\n")
	} else {
		fmt.Printf("Can't make it\n")
	}
}

func canMakeItModified(e engine, miles uint8) {
	if miles <= e.milesLeft(){
		fmt.Printf("Can make it\n")
	} else {
		fmt.Printf("Can't make it\n")
	}
}


func main() {
	var zeroEngine gasEngine = gasEngine{} // Zero Struct with default values
	var myEngine gasEngine = gasEngine{mpg: 25, gallons: 89}
	var myElectricEngine electricEngine = electricEngine{mpkwh: 34, kwh: 120}
	fmt.Println(zeroEngine, myEngine.mpg, myEngine.gallons, myElectricEngine)

	// Can do this but not reusable
	// var Engine2 = struct{
	// 	mpg uint8
	// 	gallons uint8
	// }{25,15}
	// fmt.Println(Engine2)

	fmt.Printf("Miles left: %v\n", myEngine.milesLeft())

	canMakeIt(myEngine, 50)
	canMakeItModified(myEngine, 60)
	canMakeItModified(myElectricEngine, 200)
}