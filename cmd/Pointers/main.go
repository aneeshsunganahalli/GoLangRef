package main

import "fmt"

func main() {
	var p *int32 = new(int32)
	var i int32
	fmt.Print(p)
	fmt.Printf("\nP points to %v\n", *p)
	fmt.Printf("Value at i: %v", i)

	// *p = 10
	// fmt.Printf("\nP points to %v\n", *p)
	// fmt.Printf("Value at i: %v", i)

	// p = &i
	// fmt.Printf("\nP points to %v\n", *p)
	// fmt.Printf("Value at i: %v", i)

	var thing1 = [5]float64{1,2,3,4,5}
	fmt.Printf("\nThe memory location of thing1 is %p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nThe result is: %v", result)
	fmt.Printf("\nThe value of thing1: %v", thing1)


}
// Passing in by value creates a new instance in memory -> Not Efficient
// func square(thing2 [5]float64) [5]float64 {		
// 	fmt.Printf("\nMemory loc of thing2: %p", &thing2)
// 	for i := range thing2 {
// 		thing2[i] = thing2[i] * thing2[i];
// 	}
// 	return thing2

func square(thing2 *[5]float64) [5]float64 {		// Modifies given array as memory loc is specified
	fmt.Printf("\nMemory loc of thing2: %p", thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i];
	}
	return *thing2
}